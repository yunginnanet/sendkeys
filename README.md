# sendkeys
[![GoDoc](https://godoc.org/github.com/yunginnanet/sendkeys?status.svg)](https://godoc.org/github.com/yunginnanet/sendkeys)
[![Go Report Card](https://goreportcard.com/badge/github.com/yunginnanet/sendkeys)](https://goreportcard.com/report/github.com/yunginnanet/sendkeys)

sendkeys is a golang module that strives to be a usability wrapper for the  [keybd_event](github.com/micmonay/keybd_event) library.

### Status

sendkeys is in early development. tests pass on a real machine, but I'm done trying to make github actions work for this one. Here's a local test output:

```
=== RUN   Test_strToKeys
    sendkeys_test.go:21: string: yeet, keys: []int{21, 18, 18, 20}
    sendkeys_test.go:21: string: YEET, keys: []int{-21, -18, -18, -20}
    sendkeys_test.go:21: string: YeeT, keys: []int{-21, 18, 18, -20}
--- PASS: Test_strToKeys (6.00s)
=== RUN   Test_NewKBWrapWithOptions
    sendkeys_test.go:87: sleeping for 250ms...
    sendkeys_test.go:31: [listener] go listenForKeys() start
    sendkeys_test.go:72: [receiver] go func() start
    sendkeys_test.go:48: Key pressed: y
    sendkeys_test.go:48: Key pressed: e
    sendkeys_test.go:48: Key pressed: e
    sendkeys_test.go:48: Key pressed: t
    sendkeys_test.go:81: [receiver] go func() return
    sendkeys_test.go:116: got 4 characters, got yeet string.
--- PASS: Test_NewKBWrapWithOptions (2.29s)
PASS
ok  	sendkeys	8.371s
```

### Compatibility

sendkeys has only been tested in Linux so far.

### Credits
*  ##### [micmonay](https://github.com/micmonay) of course, for creating [keybd_event](github.com/micmonay/keybd_event).
* ##### [Christopher Latham Sholes](https://en.wikipedia.org/wiki/Christopher_Latham_Sholes) for his work on the QWERTY keyboard.
