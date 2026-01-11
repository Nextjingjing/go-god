package main

import (
	"github.com/Nextjingjing/go-god/hello"
	"github.com/Nextjingjing/go-god/parent/child"
)

func main() {
	println("Hello, main!")

	hello.Greet()
	// This will cause a compile-time error
	// hello.privateFunction()

	hello.PublicFunction()
	hello.ExportedPrivateFunction()
	println(hello.GetPackageVar())

	child.ChildFunc()
}
