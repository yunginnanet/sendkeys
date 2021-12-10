# sendkeys
[![GoDoc](https://godoc.org/git.tcp.direct/kayos/sendkeys?status.svg)](https://godoc.org/git.tcp.direct/kayos/sendkeys)
[![Go Report Card](https://goreportcard.com/badge/github.com/yunginnanet/sendkeys)](https://goreportcard.com/report/github.com/yunginnanet/sendkeys)

## Summary

  sendkeys is a cross-platform usability wrapper for the [keybd_event](https://github.com/micmonay/keybd_event) Go library. It aims to to provide a *faster*, *more reliable*, and *easier to use* way for humans to access the functionality of [keybd_event](https://github.com/micmonay/keybd_event).

  Use this library to turn full strings into simulated keyboard events with ease. This library was created after a noVNC instance I was using had a broken clipboard feature.

  I have successfully used the [example](./_example/main.go) to send a very long password into a NoVNC instance that had all kinds of varying case alphanumeric characters along with many symbols.

## Features

* Optionally randomized delays between keypresses.

* Optimized map lookups should provide very high performance.

* Negative integer -> abs inversion to determine when to send the shift key event.

* Only send one key at a time, and clear the state inbetween keys for reliable functionality.

## Documentation

#### For simple usage, take a look at [the example](./_example/main.go).

<details>
  <summary>GoDoc</summary>

#### type KBOpt

```go
type KBOpt uint8
```

KBOpt[s] are options for our wrapper

```go
const (
	// Stubborn will cause our sequences to continue despite errors.
	// Otherwise, we will stop if our error count is over 0.
	Stubborn KBOpt = iota
	// Noisy will cause all errors to be printed to stdout.
	Noisy
	// Random will use random sleeps throughout the typing process.
	// Otherwise, a static 10 milliseconds will be used.
	Random
	// NoDelay will bypass the 2 second delay for linux, mostly for testing.
	NoDelay
)
```

#### type KBWrap

```go
type KBWrap struct {
	// There are unexported fields
}
```

KBWrap is a wrapper for the keybd_event library for convenience

#### func  NewKBWrapWithOptions

```go
func NewKBWrapWithOptions(opts ...KBOpt) (kbw *KBWrap, err error)
```
NewKBWrapWithOptions creates a new keyboard wrapper with the given options. As
of writing, those options include: Stubborn Noisy and Random. The defaults are
all false.

#### func (*KBWrap) BackSpace

```go
func (kb *KBWrap) BackSpace()
```
BackSpace presses the backspace key. All other keys will be cleared.

#### func (*KBWrap) Enter

```go
func (kb *KBWrap) Enter()
```
Enter presses the enter key. All other keys will be cleared.

#### func (*KBWrap) Escape

```go
func (kb *KBWrap) Escape()
```
Escape presses the escape key. All other keys will be cleared.

#### func (*KBWrap) Tab

```go
func (kb *KBWrap) Tab()
```
Tab presses the tab key. All other keys will be cleared.

#### func (*KBWrap) Type

```go
func (kb *KBWrap) Type(s string) error
```
Type types out a string by simulating keystrokes. Check the exported Symbol map
for non-alphanumeric keys.

</details>

## Status

sendkeys is in early development. tests pass on a real machine, but I'm done trying to make github actions work for this one.

<a name="test">
</a><details>
  <summary>Local Test Results</summary>

```
=== RUN   Test_strToKeys
    sendkeys_test.go:51: string: yeet, keys: []int{21, 18, 18, 20}
    sendkeys_test.go:51: string: YEET, keys: []int{-21, -18, -18, -20}
    sendkeys_test.go:51: string: YeeT, keys: []int{-21, 18, 18, -20}
--- PASS: Test_strToKeys (0.00s)
=== RUN   Test_NewKBWrapWithOptions
    sendkeys_test.go:72: [OPT] Noisy: true NoDelay: true Stubborn: true Random: true
    sendkeys_test.go:86: [OPT] Noisy: false NoDelay: false Stubborn: false Random: false
--- PASS: Test_NewKBWrapWithOptions (2.00s)
=== RUN   Test_sendkeys
    sendkeys_test.go:27: Key pressed: y
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: t
    sendkeys_test.go:171: got 4 characters: yeet
    sendkeys_test.go:27: Key pressed: Y
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: T
    sendkeys_test.go:171: got 4 characters: YeeT
    sendkeys_test.go:27: Key pressed: Y
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: t
    sendkeys_test.go:27: Key pressed: !
    sendkeys_test.go:171: got 5 characters: Yeet!
    sendkeys_test.go:27: Key pressed: \
    sendkeys_test.go:27: Key pressed: '
    sendkeys_test.go:27: Key pressed: `
    sendkeys_test.go:27: Key pressed: /
    sendkeys_test.go:27: Key pressed: 3
    sendkeys_test.go:27: Key pressed: 3
    sendkeys_test.go:27: Key pressed: 7
    sendkeys_test.go:27: Key pressed: !
    sendkeys_test.go:27: Key pressed: '
    sendkeys_test.go:27: Key pressed: \
    sendkeys_test.go:171: got 10 characters: \'`/337!'\
    sendkeys_test.go:27: Key pressed: W
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: l
    sendkeys_test.go:27: Key pressed: c
    sendkeys_test.go:27: Key pressed: o
    sendkeys_test.go:27: Key pressed: m
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:27: Key pressed: t
    sendkeys_test.go:27: Key pressed: o
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:27: Key pressed: y
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: t
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:27: Key pressed: t
    sendkeys_test.go:27: Key pressed: o
    sendkeys_test.go:27: Key pressed: w
    sendkeys_test.go:27: Key pressed: n
    sendkeys_test.go:27: Key pressed: ,
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:27: Key pressed: b
    sendkeys_test.go:27: Key pressed: u
    sendkeys_test.go:27: Key pressed: d
    sendkeys_test.go:27: Key pressed: d
    sendkeys_test.go:27: Key pressed: y
    sendkeys_test.go:27: Key pressed: !
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:171: got 28 characters: Welcome to yeet town, buddy!
    sendkeys_test.go:27: Key pressed: `
    sendkeys_test.go:27: Key pressed: ~
    sendkeys_test.go:27: Key pressed: !
    sendkeys_test.go:27: Key pressed: @
    sendkeys_test.go:27: Key pressed: #
    sendkeys_test.go:27: Key pressed: $
    sendkeys_test.go:27: Key pressed: %
    sendkeys_test.go:27: Key pressed: ^
    sendkeys_test.go:27: Key pressed: &
    sendkeys_test.go:27: Key pressed: *
    sendkeys_test.go:27: Key pressed: (
    sendkeys_test.go:27: Key pressed: )
    sendkeys_test.go:27: Key pressed: -
    sendkeys_test.go:27: Key pressed: _
    sendkeys_test.go:27: Key pressed: =
    sendkeys_test.go:27: Key pressed: +
    sendkeys_test.go:27: Key pressed: '
    sendkeys_test.go:27: Key pressed: ;
    sendkeys_test.go:27: Key pressed: :
    sendkeys_test.go:27: Key pressed: <
    sendkeys_test.go:27: Key pressed: >
    sendkeys_test.go:27: Key pressed: /
    sendkeys_test.go:27: Key pressed: \
    sendkeys_test.go:27: Key pressed: ,
    sendkeys_test.go:27: Key pressed: .
    sendkeys_test.go:27: Key pressed: |
    sendkeys_test.go:27: Key pressed: {
    sendkeys_test.go:27: Key pressed: }
    sendkeys_test.go:27: Key pressed: [
    sendkeys_test.go:27: Key pressed: ]
    sendkeys_test.go:27: Key pressed: `
    sendkeys_test.go:27: Key pressed: ~
    sendkeys_test.go:27: Key pressed: ,
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:27: Key pressed: y
    sendkeys_test.go:27: Key pressed: o
    sendkeys_test.go:27: Key pressed: u
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:27: Key pressed: f
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed: l
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:27: Key pressed: m
    sendkeys_test.go:27: Key pressed: e
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:27: Key pressed: d
    sendkeys_test.go:27: Key pressed: a
    sendkeys_test.go:27: Key pressed: w
    sendkeys_test.go:27: Key pressed: g
    sendkeys_test.go:27: Key pressed: ?
    sendkeys_test.go:27: Key pressed:
    sendkeys_test.go:32: spacebar detected
    sendkeys_test.go:171: got 52 characters: `~!@#$%^&*()-_=+';:<>/\,.|{}[]`~, you feel me dawg?
--- PASS: Test_sendkeys (6.09s)
PASS
ok  	git.tcp.direct/kayos/sendkeys	8.139s

```

</details>

## Compatibility

~~sendkeys has only been tested in Linux so far~~

the underlying library seemingly has support for all Go platforms. This should be cross platform.

Recently briefly tested in windows, I'm not sure that the shift key trigger is working or not. Needs to be tested further.

## Credits

*  ##### [micmonay](https://github.com/micmonay) of course, for creating [keybd_event](https://github.com/micmonay/keybd_event).
* ##### [Christopher Latham Sholes](https://en.wikipedia.org/wiki/Christopher_Latham_Sholes) for his work on the QWERTY keyboard.
