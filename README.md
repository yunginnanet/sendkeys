# sendkeys
[![GoDoc](https://godoc.org/git.tcp.direct/kayos/sendkeys?status.svg)](https://godoc.org/git.tcp.direct/kayos/sendkeys)
[![Go Report Card](https://goreportcard.com/badge/github.com/yunginnanet/sendkeys)](https://goreportcard.com/report/github.com/yunginnanet/sendkeys)

### sendkeys is a (*probably*) cross-platform Go usability wrapper for the  [keybd_event](https://github.com/micmonay/keybd_event) library.
  
  ---
    
      
 ### Summary     
  Use this library to turn full strings into simulated keyboard events with ease;   
  along with some neat features like with **optionally randomized delays**.   
      
 ---
        
### Improvements   
* **optimized map lookups**  

* **negative integer -> abs inversion to determine shift key necessity**  

  
  
* **It only sends one key at a time**

#### This appears to provide a faster, safer, more reliable, and easier to use alternative to [keybd_event](https://github.com/micmonay/keybd_event).
  

these tricks should should result in a very fast and clean implementation of [keybd_event](https://github.com/micmonay/keybd_event).  
   
#### See [my test results](#test) below and compare this wrapper with trying to use [keybd_event](https://github.com/micmonay/keybd_event) alone.
  
---
    
![GoDoc image](https://tcp.ac/i/baROs)

### Usage  
  
#### Check the [unit test](./sendkeys_test.go) and the [docs](https://godoc.org/git.tcp.direct/kayos/sendkeys).

### Status

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

### Compatibility

sendkeys has only been tested in Linux so far, however the underlying library seemingly has support for all Go platforms. This should be cross platform.

### Credits
*  ##### [micmonay](https://github.com/micmonay) of course, for creating [keybd_event](https://github.com/micmonay/keybd_event).
* ##### [Christopher Latham Sholes](https://en.wikipedia.org/wiki/Christopher_Latham_Sholes) for his work on the QWERTY keyboard.
