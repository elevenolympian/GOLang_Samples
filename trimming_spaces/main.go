package main 

import (
	"fmt"
	"strings"
)

func main() {
	helloWorldText := "\t Hello World"
	fmt.Printf("%d %s \n", len(helloWorldText), helloWorldText)

	trimmed := strings.TrimSpace(helloWorldText)
	fmt.Printf("%d %s \n", len(trimmed), trimmed)
}