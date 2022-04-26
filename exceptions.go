package exceptions

type try struct {
	do func()
}

type catchOrFinally struct {
	catch   func(any)
	finally func()
}

func Do(try try, catchesOrFinally ...catchOrFinally) {
	var catches []func(any)
	var finally func()
	for _, cof := range catchesOrFinally {
		if cof.catch != nil {
			catches = append(catches, cof.catch)
		}
		if cof.finally != nil {
			finally = cof.finally
			break // Do not respect catches that come after a call to finally
		}
	}

	defer func() {
		if finally != nil {
			finally()
		}

		if exc := recover(); exc != nil {
			for _, catch := range catches {
				catch(exc)
			}
		}
	}()

	try.do()
}

func Try(do func()) try {
	if do == nil {
		do = func() {}
	}
	return try{do: do}
}

func Catch[C any](do func(C)) catchOrFinally {
	if do == nil {
		do = func(C) {}
	}
	return catchOrFinally{
		catch: func(caught any) {
			if c, ok := caught.(C); ok {
				do(c)
			}
		},
	}
}

func Finally(do func()) catchOrFinally {
	if do == nil {
		do = func() {}
	}
	return catchOrFinally{
		finally: do,
	}
}
