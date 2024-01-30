package main

import (
	"fmt"
	"gopl/chapter2/tempconv0"
)

func main() {
	fmt.Printf("%g\n", tempconv0.BoilingC-tempconv0.FreezingC)
	boilingF := tempconv0.CToF(tempconv0.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv0.CToF(tempconv0.FreezingC))
	c := tempconv0.FToC(212.0)
	fmt.Println(c.String())
	var k tempconv0.Kelvin
	fmt.Printf("Kelvin: %s\tCelsius: %s\t Fahrenheit: %s\n", k.String(), tempconv0.KToC(k).String(), tempconv0.KToF(k).String())
}
