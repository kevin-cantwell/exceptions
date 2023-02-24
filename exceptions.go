package exceptions

func Try(do func(), thens ...then) {
	defer func() {
		for _, then := range thens {
			if then.finally != nil {
				then.finally()
			}
		}

		if cause := recover(); cause != nil {
			for _, then := range thens {
				if then.catch != nil {
					then.catch(cause)
				}
			}
		}
	}()

	do()
}

type then struct {
	catch   func(any)
	finally func()
}

func Catch[C any](do func(C)) then {
	if do == nil {
		do = func(C) {}
	}
	return then{
		catch: func(cause any) {
			if c, ok := cause.(C); ok {
				do(c)
			}
		},
	}
}

func Finally(do func()) then {
	if do == nil {
		do = func() {}
	}
	return then{
		finally: do,
	}
}
