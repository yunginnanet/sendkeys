package sendkeys

import (
	"strings"
	"testing"
	"time"

	"github.com/eiannone/keyboard"
)

func strTo(teststr string, t *testing.T) {
	split := strings.Split(teststr, "")
	k, err := NewKBWrapWithOptions(Noisy)
	if err != nil {
		t.Fatalf(err.Error())
	}
	keys := k.strToKeys(teststr)
	if len(keys) != len(split) {
		t.Fatalf("length of mapped keys: %d, wanted length of string: %d", len(keys), len(split))
	}
	t.Logf("string: %s, keys: %#v", teststr, keys)
}

func Test_strToKeys(t *testing.T) {
	strTo("yeet", t)
	strTo("YEET", t)
	strTo("YeeT", t)
}

func listenForKeys(t *testing.T, ret chan string) {
	t.Log("[listener] go listenForKeys() start")
	defer t.Log("[listener go listenForKeys() return")

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		_ = keyboard.Close()
	})

	for {
		event := <-keysEvents
		if event.Err != nil {
			t.Fatalf(event.Err.Error())
		}
		t.Logf("Key pressed: %s", string(event.Rune))
		if event.Key == keyboard.KeyEsc {
			break
		}
		ret <- string(event.Rune)
	}
}

func Test_NewKBWrapWithOptions(t *testing.T) {
	var teststr = "yeet"

	k, err := NewKBWrapWithOptions(Noisy)
	if err != nil {
		t.Fatal(err)
	}
	keys := k.strToKeys(teststr)
	ret := make(chan string, len(teststr))
	go listenForKeys(t, ret)
	k.set(keys...)

	var count = 0
	var chars []string

	go func() {
		t.Log("[receiver] go func() start")
		defer t.Log("[receiver] go func() return")
		for {
			select {
			case chr := <-ret:
				chars = append(chars, chr)
				count++
			default:
				if count == len(teststr) {
					return
				}
			}
		}
	}()

	t.Log("sleeping for 250ms...")
	time.Sleep(250 * time.Millisecond)

	err = k.Type(teststr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	var tick = 0
	var brk = false
	for {
		if brk {
			break
		}
		switch {
		case tick >= 5:
			t.Logf("FAIL: took too long")
			t.Fail()
			brk = true
			break
		case count == len(teststr):
			brk = true
			break
		default:
			time.Sleep(1 * time.Second)
			tick++
		}
	}

	t.Logf(
		"got %d characters, got %s string.",
		count, strings.Join(chars, ""),
	)

	if strings.Join(chars, "") != teststr {
		t.Fail()
	}
}
