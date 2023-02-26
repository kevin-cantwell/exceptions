This package is bad and you shouldn't use it.

Seriously, don't even read this.

Turn back now.

Still here? Ok cool. 

Package `exceptions` brings the old try/catch/finally pattern you know and love to Go. It's a complete abomination of everything that is holy about Go and if you use it your peers will mock you and your friends will abandon you. Here's how it works:

https://github.com/kevin-cantwell/exceptions/blob/e981e831ad71acb2ef102cda772d983e5dbd2bd6/_examples/gofishing/main.go#L10-L26

The above program outputs:
```
Trying to fish...
I caught a Shark!
Stopped fishing.
```

https://github.com/kevin-cantwell/exceptions/blob/e981e831ad71acb2ef102cda772d983e5dbd2bd6/_examples/runtime_error/main.go#L10-L20

The above program outputs:
```
caught: runtime error: index out of range [2] with length 0
This will print after catching the error.
```

https://github.com/kevin-cantwell/exceptions/blob/28e7d3e02e0e5c0daba6709607616958cb1c94a0/_examples/nilpanic/main.go#L10-L19

The above program outputs:
```
nil panic! <nil>
This will print after catching the error.
```

This follows the semantics of try...catch...finally where any (optional) `Finally` funcs are invoked directly after `Try` any `Catch` that was invoked. The `Catch` functions take a single arguments: a "cause" of any type. The cause is any type you might expect to be recovered from a panic. Only the first catch that meets these requirements is invoked.

Ironically, this package declares no exception types.

Requires go version 1.18 or greater (for generics)
