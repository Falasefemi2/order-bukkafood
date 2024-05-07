package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func displayGeneratingBill() {
	fmt.Println()
	billDisplayText := "************************************* Generating Bill *************************************"
	for _, txt := range billDisplayText {
		fmt.Printf("%c", txt)
		time.Sleep(time.Millisecond * 15)
	}
}

func generateBill() {
	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf(" %-30s %-20s %-20s %-20s\n", "Food Name", "Price(₦)", "Quantity", "Total Price(₦)")
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	printBill()

	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	//print sub total amount in a cool way!
	fmt.Printf("%47s", " ")
	for _, element := range "Sub Total(excluding GST): ₦" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("%.2f\n", subTotalBill)
}

// prints the bill
func printBill() {
	for key := range customerOrder {
		for _, value := range menu {
			if key == value.itemName {
				totalCostFood := float64(customerOrder[key]) + value.itemPrice
				fmt.Printf(" %-30s %-20.2f %-20v %-20.2f\n", key, value.itemPrice, customerOrder[key], totalCostFood)

			}
		}
	}
	fmt.Println()
}

func printFinalBill() {
	for _, element := range "Here is the final bill-" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println()

	fmt.Printf("\n%52s\n", "FALASE FEMI")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s\n", strings.Repeat("*", 91))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%86s\n", "37 Oba Ni Jesu Street, Ijoko, Ogun State")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%50s\n", "Tel: +2347013329953")
	fmt.Printf("%60s\n\n", "Email: femifalase228@gmail.com")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s", strings.Repeat("-", 42))
	fmt.Printf("%s", "INVOICE")
	fmt.Printf("%s\n", strings.Repeat("-", 42))
	time.Sleep(time.Millisecond * 200)

	randSource := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(randSource)

	randGenerator.Intn(100) //necessary to produce random integers
	fmt.Printf(" Ticket No: %d\n", rand.Intn(550)+1)

	fmt.Printf(" Date: %v\n", time.Now().Local().Format("06-Jan-02")) //display date
	fmt.Printf(" Time: %v", time.Now().Local().Format("3:4 PM"))      //display time
	time.Sleep(time.Millisecond * 200)

	generateBill() //prints details of the bill,like, item name, price, quantity and total price and sub total amount.

	tax := 18 * subTotalBill / (100)
	grandTotal := subTotalBill + tax

	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: $%.2f\n", "GST", tax) //display tax amount
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: $%.2f\n", "Grand Total", grandTotal) //display final bill
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
}
