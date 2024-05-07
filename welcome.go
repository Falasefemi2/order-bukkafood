package main

import "fmt"

func welcome(customerName string) {
	fmt.Printf("%58s %s\n", "Welcome ", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to Bukka Hut! _/\\_ ")
	fmt.Println()
}

func bye(customerName string) {
	fmt.Println()
	fmt.Printf("%72s\n", "_/\\_ Goodbye, Enjoy your meal! _/\\_ ")
	fmt.Printf("%60s %s\n", "We hope to see you again!", customerName)
}
