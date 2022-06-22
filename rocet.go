package main

import "fmt"

func main() {
	const hoursPerDay = 24

	var days = 28
	var distance = 56000000000S0 // km

	fmt.Println(distance/(days*hoursPerDay), "км/ч")
}
