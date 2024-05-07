package main

import (
	"fmt"
	"strings"
)

func order() {
	printMenu()
	var foodNumber int
	var noOfPlate int
	var foodName string
	var foodPrice float64
	var choiceName string

	for {
		fmt.Println()
		fmt.Println("Enter 0 to exit...")
		fmt.Print("Enter food number of food to order: ")
		fmt.Scan(&foodNumber)
		if foodNumber == 0 {
			break
		}

		for index, item := range menu {
			if index+1 == foodNumber {
				foodName = item.itemName
				foodPrice = item.itemPrice
				choiceName = item.itemName               // Assign choiceName here
				fmt.Printf("You ordered %s\n", foodName) // Print the food name
				break
			}
		}
		fmt.Printf("How many %v plates do you want: ", choiceName)
		fmt.Scan(&noOfPlate)
		if noOfPlate == 0 {
			continue
		} else {
			for index := range menu {
				if index+1 == foodNumber {
					customerOrder[choiceName] += noOfPlate
					subTotalBill += foodPrice * float64(noOfPlate)
					break
				}
			}
			fmt.Printf("\nYou just ordered %v %v which amounts to ₦%v.\n", noOfPlate, choiceName, foodPrice*float64(noOfPlate))
			orderTillNow()
		}
	}
}

// print everytime you order a new food
func orderTillNow() {
	// print what you ordered now
	fmt.Println("\nYour order till now")
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Printf(" %-12s %s\n", "Quantity", "Item")
	fmt.Printf("%s\n", strings.Repeat("-", 32))

	for order := range customerOrder {
		fmt.Printf(" %-12d %s\n", customerOrder[order], order)
	}
	fmt.Printf("%s\n", strings.Repeat("-", 32))

}
