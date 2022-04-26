package exceptions_test

import (
	"testing"

	"github.com/kevin-cantwell/exceptions"
)

type throwable1 struct {
	Cause string
}

type throwable2 struct {
	Cause string
}

func TestExceptions(t *testing.T) {
	exceptions.Do(
		exceptions.Try(func() {
			panic("foo")
		}),
		exceptions.Catch(func(e string) {
			t.Log("caught string:", e)
		}),
		exceptions.Catch(func(e int64) {
			t.Log("caught int64:", e)
		}),
		exceptions.Finally(func() {
			t.Log("finally!")
		}),
	)
	t.FailNow()
}

// var _ = Describe("#Tryer", func() {
// 	It("Should try, catch, finally", func() {
// 		var exception1 throwable1
// 		var exception2 throwable2
// 		var finally string

// 		exceptions.Try(func() {
// 			finally = "try"
// 			panic(throwable2{Cause: "panic!"})
// 		}).Catch[throwable1](func() {
// 			Fail("Should not catch this type")
// 		}).Catch(&exception2, func() {
// 			Expect(exception2.Cause).To(Equal("panic!"))
// 			Expect(finally).To(Equal("finally"))
// 		}).Finally(func() {
// 			Expect(finally).To(Equal("try"))
// 			finally = "finally"
// 		}).Do()

// 		Expect(exception2.Cause).To(Equal("panic!"))
// 	})
// })
