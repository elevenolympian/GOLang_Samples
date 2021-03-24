package main

import "fmt"

func main() {
	const freezingF, boilingF, randomF = 32.0, 212.0 , 68.4
	fmt.Printf("%g Fahrenheit = %g Centigrad \n", freezingF, FahrenheitToCelcius(freezingF))
	fmt.Printf("%g Fahrenheit = %g Centigrad \n", boilingF, FahrenheitToCelcius(boilingF))
	fmt.Printf("%g Random Fahrenheit = %g Centigrad \n", randomF, FahrenheitToCelcius(randomF))
}

func FahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9;
}