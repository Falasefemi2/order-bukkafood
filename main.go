package main

import "fmt"

var subTotalBill float64

var customerOrder = make(map[string]int, 0)

func main() {
	var customerName string
	fmt.Println("What is your name? ")
	fmt.Scan(&customerName)

	welcome(customerName)
	order()
	displayGeneratingBill()
	generateBill()

	saveRecipt()
	printFinalBill()

	bye(customerName)
}
