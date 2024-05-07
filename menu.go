package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    int
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	{1, "Jollof Rice", 20.88},
	{2, "Fried Rice", 18.99},
	{3, "Chicken Curry", 22.50},
	{4, "Beef Stew", 25.75},
	{5, "Vegetable Fried Noodles", 16.20},
	{6, "Grilled Salmon", 28.95},
	{7, "Margherita Pizza", 14.50},
	{8, "Spaghetti Bolognese", 19.75},
	{9, "Greek Salad", 12.80},
	{10, "Beef Burger", 16.95},
	{11, "Vegetable Lasagna", 18.25},
	{12, "Pad Thai", 17.50},
	{13, "Chicken Fajitas", 21.30},
	{14, "Caesar Salad", 13.95},
	{15, "Mushroom Risotto", 20.60},
}

func printMenu() {
	fmt.Printf("%-55s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	fmt.Printf("%-5s %-30s %10s\n", "S.No", "Item Name", "Price")
	fmt.Printf("%s\n", strings.Repeat("-", 50))
	for _, food := range menu {
		fmt.Printf("%-5d %-30s ₦%7.2f\n", food.itemNo, food.itemName, food.itemPrice)
	}
	fmt.Printf("%s\n", strings.Repeat("-", 50))
}
