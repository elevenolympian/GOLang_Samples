package hello

import "rsc.io/quote"

func SimpleFunction(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func SayHelloFromtheModule() string {
	return "Hello from the module"
}

func Hello() string {
	return "Hello World String"
}

func SampleString() string {
	return "sample Test case"
}

func TakeAQuote() string {
	return quote.Hello()
}
