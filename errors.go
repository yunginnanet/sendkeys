package sendkeys

import "errors"

// ErrKeyMappingNotFound is an error returned when we don't know how to handle the given character.
var ErrKeyMappingNotFound = errors.New("failed to map key: ")

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
	if err == nil {
		return
	}
	kb.errors = append(kb.errors, err)
	if kb.noisy {
		println(err.Error())
	}
}
