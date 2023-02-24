This package is bad and you shouldn't use it.

Seriously, don't even read this.

Turn back now.

Still here? Ok cool. 

Package `exceptions` brings the old try/catch/finally pattern you know and love to Go. It's a complete abomination of everything that is holy about Go and if you use it your peers will mock you and your friends will abandon you. Here's how it works:

https://github.com/kevin-cantwell/exceptions/blob/0dbc3f451ae37674bcbd4907c00e3be67acd1d18/_examples/gofishing/main.go#L1-L32

The above program outputs:

```
Trying to fish...
Stopped fishing.
I caught a Shark! Vegetarian?: true
```

This follows the semantics of try...catch...finally where the (optional) `Finally` func is invoked directly after `Try` and before any `Catch` function is called. The `Catch` function takes two arguments: an "exception", and a function. The exception is a pointer to a struct and is set if and only if the panic value is of the same indirect type. Only the first catch that meets these requirements is invoked.

Ironically, this package declares no exception types.

Requires go version 1.18 or greater (for generics)
