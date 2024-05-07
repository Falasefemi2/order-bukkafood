package main

import (
	"fmt"
	"io"
	"os"
)

func saveRecipt() {
	file, err := os.CreateTemp("", "bill*.txt")
	if err != nil {
		panic(err)
	}

	defer os.Remove(file.Name())

	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()
	os.Stdout = file
	printFinalBill()

	file.Seek(0, 0)
	billBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading temp file:", err)
		return
	}
	billString := string(billBytes)

	// Save the bill to the final file
	saveBillToFile(billString, customerName)
}

func saveBillToFile(bill string, customerName string) {
	fileName := customerName + "_bill.txt"
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(bill)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Printf("Bill saved to %s\n", fileName)
}
