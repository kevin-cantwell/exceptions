package exceptions

import "reflect"

// Tryer provides the semantics of try...catch...finally
type Tryer struct {
	try     func()
	excptns []interface{}
	catches []func()
	finally func()
}

// Try accepts a func to try and returns a pointer to Tryer. Caller must invoke
// Do to execute try.
func Try(try func()) *Tryer {
	return &Tryer{try: try}
}

// Catch registers an exception and a corresponding function with Tryer for matching against
// recovered types. The exception must be a pointer type of the expected recovered type.
func (try *Tryer) Catch(exception interface{}, catch func()) *Tryer {
	try.excptns = append(try.excptns, exception)
	try.catches = append(try.catches, catch)
	return try
}

// Finally registers a function to be executed directly after Try
func (try *Tryer) Finally(finally func()) *Tryer {
	try.finally = finally
	return try
}

// Do begins the execution of the try/catch/finally system. The try function
// will be called, followed by finally. If a panic occurs during try, it will
// be recovered and the resulting value will be compared against each registered catch
// in order of registration. If a recovered type can be assigned to a catch's exception it
// will be assigned and its corresponding function invoked.
func (try *Tryer) Do() {
	defer func() {
		try.finally()

		if ex := recover(); ex != nil {
			for i := range try.excptns {
				if reflect.TypeOf(try.excptns[i]) == reflect.PtrTo(reflect.TypeOf(ex)) {
					value := reflect.ValueOf(try.excptns[i])
					value.Elem().Set(reflect.ValueOf(ex))
					try.catches[i]()
					return
				}
			}
		}
	}()
	try.try()
}

// Throw exists for no good reason. All it does is call panic.
func Throw(t interface{}) {
	panic(t)
}
