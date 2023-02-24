This package is bad and you shouldn't use it.

Seriously, don't even read this.

Turn back now.

Still here? Ok cool. 

Package `exceptions` brings the old try/catch/finally pattern you know and love to Go. It's a complete abomination of everything that is holy about Go and if you use it your peers will mock you and your friends will abandon you. Here's how it works:

https://github.com/kevin-cantwell/exceptions/blob/f50a8a3c36260ad3b99f9e97c8c619304daecf20/_examples/gofishing/main.go#L9-L27

The above program outputs:

```
Trying to fish...
Stopped fishing.
I caught a Shark!
```

This follows the semantics of try...catch...finally where any (optional) `Finally` funcs are invoked directly after `Try` and before any `Catch` function is called. The `Catch` functions take a single arguments: a "cause" of any type. The cause is any type you might expect to be recovered from a panic. Only the first catch that meets these requirements is invoked.

Ironically, this package declares no exception types.

Requires go version 1.18 or greater (for generics)
