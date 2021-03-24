package main 

import (
	"strings"
	"fmt"
)

func main() {
	helloWorld := "We are good! Thanks for this to the World."
	helloMars := strings.Replace(helloWorld, "World", "Mars", -1)
	fmt.Println(helloMars)
}