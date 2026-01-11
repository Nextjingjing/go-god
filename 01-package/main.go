package main

import (
	"github.com/Nextjingjing/01-package/go-god/hello"
	"github.com/Nextjingjing/01-package/go-god/parent/child"
)

func main() {
	println("Hello, main!")

	hello.Greet()
	hello.Greet2()
	// This will cause a compile-time error
	// hello.privateFunction()

	hello.PublicFunction()
	hello.ExportedPrivateFunction()
	println(hello.GetPackageVar())

	child.ChildFunc()
}
