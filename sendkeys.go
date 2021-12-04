package sendkeys

import (
	"errors"
	"time"

	"github.com/micmonay/keybd_event"
)

// KBWrap is a wrapper for the keybd_event library for convenience
type KBWrap struct {
	backend  keybd_event.KeyBonding
	errors   []error
	stubborn bool
	noisy    bool
	random   bool
}


func newKbw() *KBWrap {
	return &KBWrap{
		errors:   []error{},
		stubborn: false,
		noisy:    false,
		random:   false,
	}
}

// NewKBWrapWithOptions creates a new keyboard wrapper with the given options.
// As of writing, those options include: Stubborn Noisy and Random.
// The defaults are all false.
func NewKBWrapWithOptions(opts ...KBOpt) (kbw *KBWrap, err error) {
	linDelay()
	kbw = newKbw()
	kbw.backend, err = keybd_event.NewKeyBonding()
	if err != nil {
		return nil, err
	}
	kbw.processOptions(opts...)
	return
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
// Default wait time is 10 milliseconds.
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
