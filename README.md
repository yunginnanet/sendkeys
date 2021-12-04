# sendkeys
[![GoDoc](https://godoc.org/git.tcp.direct/kayos/sendkeys?status.svg)](https://godoc.org/git.tcp.direct/kayos/sendkeys)
[![Go Report Card](https://goreportcard.com/badge/github.com/yunginnanet/sendkeys)](https://goreportcard.com/report/github.com/yunginnanet/sendkeys)

sendkeys is a golang module that strives to be a usability wrapper for the  [keybd_event](https://github.com/micmonay/keybd_event) library.

### Status

sendkeys is in early development. tests pass on a real machine, but I'm done trying to make github actions work for this one.

<details>
  <summary>Local Test Results</summary>

```
=== RUN   Test_strToKeys
    sendkeys_test.go:44: string: yeet, keys: []int{21, 18, 18, 20}
    sendkeys_test.go:44: string: YEET, keys: []int{-21, -18, -18, -20}
    sendkeys_test.go:44: string: YeeT, keys: []int{-21, 18, 18, -20}
--- PASS: Test_strToKeys (0.00s)
=== RUN   Test_NewKBWrapWithOptions
    sendkeys_test.go:65: [OPT] Noisy: true NoDelay: true Stubborn: true Random: true
    sendkeys_test.go:79: [OPT] Noisy: false NoDelay: false Stubborn: false Random: false
--- PASS: Test_NewKBWrapWithOptions (2.00s)
=== RUN   Test_sendkeys
    sendkeys_test.go:26: Key pressed: y
    sendkeys_test.go:26: Key pressed: e
    sendkeys_test.go:26: Key pressed: e
    sendkeys_test.go:26: Key pressed: t
    sendkeys_test.go:150: got 4 characters, got yeet string.
    sendkeys_test.go:26: Key pressed: Y
    sendkeys_test.go:26: Key pressed: e
    sendkeys_test.go:26: Key pressed: e
    sendkeys_test.go:26: Key pressed: T
    sendkeys_test.go:150: got 4 characters, got YeeT string.
    sendkeys_test.go:26: Key pressed: Y
    sendkeys_test.go:26: Key pressed: e
    sendkeys_test.go:26: Key pressed: e
    sendkeys_test.go:26: Key pressed: t
    sendkeys_test.go:26: Key pressed: !
    sendkeys_test.go:150: got 5 characters, got Yeet! string.
    sendkeys_test.go:26: Key pressed: $
    sendkeys_test.go:26: Key pressed: #
    sendkeys_test.go:26: Key pressed: !
    sendkeys_test.go:26: Key pressed: Y
    sendkeys_test.go:26: Key pressed: ^
    sendkeys_test.go:26: Key pressed: %
    sendkeys_test.go:26: Key pressed: #
    sendkeys_test.go:26: Key pressed: *
    sendkeys_test.go:26: Key pressed: *
    sendkeys_test.go:26: Key pressed: (
    sendkeys_test.go:26: Key pressed: #
    sendkeys_test.go:26: Key pressed: (
    sendkeys_test.go:26: Key pressed: @
    sendkeys_test.go:26: Key pressed: ^
    sendkeys_test.go:26: Key pressed: ?
    sendkeys_test.go:26: Key pressed: ?
    sendkeys_test.go:26: Key pressed: !
    sendkeys_test.go:26: Key pressed: ?
    sendkeys_test.go:26: Key pressed: !
    sendkeys_test.go:26: Key pressed: `
    sendkeys_test.go:26: Key pressed: `
    sendkeys_test.go:26: Key pressed: `
    sendkeys_test.go:26: Key pressed: `
    sendkeys_test.go:26: Key pressed: `
    sendkeys_test.go:26: Key pressed: "
    sendkeys_test.go:26: Key pressed: _
    sendkeys_test.go:26: Key pressed: _
    sendkeys_test.go:26: Key pressed: t
    sendkeys_test.go:26: Key pressed: !
    sendkeys_test.go:150: got 29 characters, got $#!Y^%#**(#(@^??!?!`````"__t! string.
    sendkeys_test.go:26: Key pressed:
--- PASS: Test_sendkeys (3.44s)
PASS
ok  	git.tcp.direct/kayos/sendkeys	5.507s
```

</details>

### Compatibility

sendkeys has only been tested in Linux so far.

### Credits
*  ##### [micmonay](https://github.com/micmonay) of course, for creating [keybd_event](https://github.com/micmonay/keybd_event).
* ##### [Christopher Latham Sholes](https://en.wikipedia.org/wiki/Christopher_Latham_Sholes) for his work on the QWERTY keyboard.
