package hello

var packageVar = "I am a package variable"

func Greet() {
	println("Hello, package!")
}

func privateFunction() {
	println("Hello, private function!")
}

func PublicFunction() {
	println("Hello, public function!")
}

func ExportedPrivateFunction() {
	privateFunction()
}

func GetPackageVar() string {
	return packageVar
}
