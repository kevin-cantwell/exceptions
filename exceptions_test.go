package exceptions_test

import (
	"github.com/kevin-cantwell/exceptions"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type throwable1 struct {
	Cause string
}

type throwable2 struct {
	Cause string
}

var _ = Describe("#Tryer", func() {
	It("Should try, catch, finally", func() {
		var exception1 throwable1
		var exception2 throwable2
		var finally string

		exceptions.Try(func() {
			finally = "try"
			panic(throwable2{Cause: "panic!"})
		}).Catch(&exception1, func() {
			Fail("Should not catch this type")
		}).Catch(&exception2, func() {
			Expect(exception2.Cause).To(Equal("panic!"))
			Expect(finally).To(Equal("finally"))
		}).Finally(func() {
			Expect(finally).To(Equal("try"))
			finally = "finally"
		}).Do()

		Expect(exception2.Cause).To(Equal("panic!"))
	})
})
