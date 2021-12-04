package sendkeys

import (
	"errors"
	"time"

	kbd "github.com/micmonay/keybd_event"
)

// KBWrap is a wrapper for the keybd_event library for convenience
type KBWrap struct {
	d        kbd.KeyBonding
	errors   []error
	stubborn bool
	noisy    bool
	random   bool
	nodelay  bool
}

func newKbw() *KBWrap {
	return &KBWrap{
		errors:   []error{},
		stubborn: false,
		noisy:    false,
		random:   false,
		nodelay:  false,
	}
}

// NewKBWrapWithOptions creates a new keyboard wrapper with the given options.
// As of writing, those options include: Stubborn Noisy and Random.
// The defaults are all false.
func NewKBWrapWithOptions(opts ...KBOpt) (kbw *KBWrap, err error) {
	kbw = newKbw()
	kbw.d, err = kbd.NewKeyBonding()
	if err != nil {
		return nil, err
	}
	kbw.processOptions(opts...)
	kbw.linDelay()
	return
}

func (kb *KBWrap) down() {
	if !kb.check() {
		return
	}
	kb.handle(kb.d.Press())
}
func (kb *KBWrap) up() {
	if !kb.check() {
		return
	}
	kb.handle(kb.d.Release())
}

// press presses a key from the queue, waits, and then releases.
// Default wait time is 10 milliseconds.
func (kb *KBWrap) press() {
	kb.down()
	if kb.random {
		snoozeMS(rng(25))
	} else {
		time.Sleep(10 * time.Millisecond)
	}
	kb.up()
}

func (kb *KBWrap) set(keys ...int) {
	kb.d.SetKeys(keys...)
}

func (kb *KBWrap) clr() {
	kb.d.Clear()
}

func (kb *KBWrap) only(k int) {
	kb.clr()
	kb.set(k)
	kb.press()
	kb.clr()
}

// Escape presses the escape key.
// All other keys will be cleared.
func (kb *KBWrap) Escape() {
	kb.only(kbd.VK_ESC)
}

// Tab presses the tab key.
// All other keys will be cleared.
func (kb *KBWrap) Tab() {
	kb.only(kbd.VK_TAB)
}

// Enter presses the enter key.
// All other keys will be cleared.
func (kb *KBWrap) Enter() {
	kb.only(kbd.VK_ENTER)
}

// BackSpace presses the backspace key.
// All other keys will be cleared.
func (kb *KBWrap) BackSpace() {
	kb.only(kbd.VK_ENTER)
}

// Type types out a string by simulating keystrokes.
// Check the exported Symbol map for non-alphanumeric keys.
func (kb *KBWrap) Type(s string) error {
	keys := kb.strToKeys(s)
	for _, key := range keys {
		if !kb.check() {
			return errors.New(compoundErr(kb.errors))
		}
		if key < 0 {
			kb.d.HasSHIFT(true)
			key = abs(key)
		}

		kb.set(key)
		kb.press()
		kb.clr()
	}
	return nil
}
