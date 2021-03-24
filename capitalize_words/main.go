package main 

import ( 
	"strings"
	"fmt"
)

func main() {
	TitleText("how are you today!")
	CapitalizedText("hello world text")
}

func TitleText(textToBeCapitalized string) string {
	titleOftheText := strings.Title(textToBeCapitalized)
	fmt.Println("entitle the text: " + titleOftheText)
	return titleOftheText
}

func CapitalizedText(paramText string) string {
	textToBeUppered := strings.ToUpper(paramText)
	fmt.Println("textToBeUppered: " + paramText)
	return textToBeCapitalized
}