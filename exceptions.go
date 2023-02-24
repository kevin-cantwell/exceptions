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

func Catch[C any](catch func(C)) then {
	if catch == nil {
		catch = func(C) {}
	}
	return then{
		catch: func(cause any) {
			if c, ok := cause.(C); ok {
				catch(c)
			}
		},
	}
}

func Finally(finally func()) then {
	if finally == nil {
		finally = func() {}
	}
	return then{
		finally: finally,
	}
}
