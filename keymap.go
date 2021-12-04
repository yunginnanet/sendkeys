package sendkeys

import (
	"errors"
	"strings"
	"time"

	"github.com/micmonay/keybd_event"
)

// KBOpt are options for our wrapper
type KBOpt uint8

// KeyMappingNotFound is an error returned when we don't know how to handle the given character.
var KeyMappingNotFound = errors.New("failed to map key")

const (
	// Stubborn will cause our sequences to continue despite errors.
	// Otherwise, we will stop if our error count is over 0.
	Stubborn KBOpt = iota
	// Noisy will cause all errors to be printed to stdout.
	Noisy
	// Random will use random sleeps throughout the typing process.
	// Otherwise, a static 10 milliseconds will be used.
	Random
)

// KBWrap is a wrapper for the keybd_event library for convenience
type KBWrap struct {
	backend  keybd_event.KeyBonding
	errors   []error
	stubborn bool
	noisy    bool
	random   bool
}

func (kb *KBWrap) processOptions(opts ...KBOpt) {
	kbo := map[KBOpt]*bool{
		Stubborn: &kb.stubborn,
		Noisy:    &kb.noisy,
		Random:   &kb.random,
	}
	for _, o := range opts {
		if option, ok := kbo[o]; !ok {
			continue
		} else {
			*option = true
		}
	}
}

func newKbw() *KBWrap {
	return &KBWrap{
		errors:   []error{},
		stubborn: false,
		noisy:    false,
		random:   false,
	}
}

func NewKBWrapWithOptions(opts ...KBOpt) (kbw *KBWrap, err error) {
	linDelay()
	kbw = newKbw()
	kbw.backend, err = keybd_event.NewKeyBonding()
	if err != nil {
		return nil, err
	}
	kbw.processOptions()
	return
}

type funcErr func() error

func (kb *KBWrap) check() bool {
	if kb.stubborn {
		return true
	}
	if len(kb.errors) > 0 {
		return false
	}
	return true
}

func (kb *KBWrap) handle(err error) {
	if err != nil {
		kb.errors = append(kb.errors, err)
		if kb.noisy {
			println(err.Error())
		}
	}
}

func (kb *KBWrap) press() {
	if !kb.check() {
		return
	}
	kb.handle(kb.backend.Press())
}
func (kb *KBWrap) release() {
	if !kb.check() {
		return
	}
	kb.handle(kb.backend.Release())
}

// pressAndRelease presses a key from the queue, waits, and then releases.
func (kb *KBWrap) pressAndRelease() {
	kb.press()
	if kb.random {
		snoozeMS(rng(25))
	} else {
		time.Sleep(10 * time.Millisecond)
	}
	kb.release()
}

func (kb *KBWrap) set(keys ...int) {
	kb.backend.SetKeys(keys...)
}

func compoundErr(errs []error) string {
	var es []string
	for _, e := range errs {
		if e == nil {
			continue
		}
		es = append(es, e.Error())
	}
	return strings.Join(es, ",")
}

// Type types out a string by simulating keystrokes.
func (kb *KBWrap) Type(s string) error {
	keys := kb.strToKeys(s)
	for _, key := range keys {
		if !kb.check() {
			return errors.New(compoundErr(kb.errors))
		}
		kb.set(key)
		kb.pressAndRelease()
	}
	return nil
}

func (kb *KBWrap) strToKeys(s string) (keys []int) {
	split := strings.Split(s, "")
	for _, c := range split {
		d, dok := num[c]
		a, aok := alpha[c]
		ca, caok := alpha[strings.ToLower(c)]

		switch {
		case aok:
			keys = append(keys, a)
		case dok:
			keys = append(keys, d)
		case caok:
			keys = append(keys, 0-ca)
		default:
			kb.errors = append(kb.errors, KeyMappingNotFound)
		}
	}
	return
}

var num = map[string]int{
	"1": 2, "2": 3, "3": 4, "4": 5, "5": 6,
	"6": 7, "7": 8, "8": 9, "9": 10, "0": 11,
}

var alpha = map[string]int{
	"q": 16, "w": 17, "e": 18, "r": 19, "t": 20,
	"y": 21, "u": 22, "i": 23, "o": 24, "p": 25,
	"a": 30, "s": 31, "d": 32, "f": 33, "g": 34,
	"h": 35, "j": 36, "k": 37, "l": 38, "z": 44,
	"x": 45, "c": 46, "v": 47, "b": 48, "n": 49,
	"m": 50,
}
