package exceptions

// Try executes the given function then invokes all Finally blocks. If the try function panics,
// the panic is recovered and its value is passed to the first Catch block that matches
// the recovered type. If no suitable Catch is found, it panics with the recovered value.
func Try(try func(), thens ...then) {
	defer func() {
		for _, then := range thens {
			if then.finally != nil {
				then.finally()
			}
		}

		if cause := recover(); cause != nil {
			for _, then := range thens {
				if then.catch != nil {
					if caught := then.catch(cause); caught {
						return
					}
				}
			}
			panic(cause)
		}
	}()

	try()
}

type then struct {
	catch   func(any) bool
	finally func()
}

// Catch calls the given function with the value recovered from a panic in the Try block.
// Catch blocks are evaluated in order the first one that matches the type of the recovered
// value is called. At most one Catch block will be called for each invocation of Try.
func Catch[C any](catch func(C)) then {
	if catch == nil {
		catch = func(C) {}
	}
	return then{
		catch: func(cause any) bool {
			if c, ok := cause.(C); ok {
				catch(c)
				return true
			}
			return false
		},
	}
}

// Finally calls the given function after the Try block completes and before any Catch block
// is run. All Finally blocks are called in the order they are given, whether the Try block
// panics or not.
func Finally(finally func()) then {
	if finally == nil {
		finally = func() {}
	}
	return then{
		finally: finally,
	}
}

// Throw panics with the given cause. Just here for OCD reasons
func Throw(cause any) {
	panic(cause)
}
