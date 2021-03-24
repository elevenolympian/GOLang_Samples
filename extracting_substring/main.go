package main 

import "fmt"

func main() {
	fmt.Println(ExtractingSubString())
}

func ExtractingSubString() string {
	textForGreetings := "Hello World text is here"
	firstPartString := textForGreetings[:5]
	secondPartString := textForGreetings[5:]
	fmt.Println(firstPartString)
	fmt.Println(secondPartString)
	return firstPartString
}