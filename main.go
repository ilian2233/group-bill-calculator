package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Bill struct {
	Name     string   `json:"name"`
	Involved []Person `json:"involved"`
}

type Person struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

var filename string

func init() {
	flag.StringVar(&filename, "file", "", "JSON file containing bills")
}

func main() {
	flag.Parse()

	if filename == "" {
		filename = "bills.json"
	}

	bills, err := readBills(filename)
	if err != nil {
		log.Fatalf("Error reading bills: %v", err)
	}

	imbalances := calculateImbalances(bills)

	if len(imbalances) == 0 {
		fmt.Println("No imbalances to correct.")
		return
	}

	fmt.Println("Final imbalances per person:")
	for name, imbalance := range imbalances {
		if imbalance < 0 {
			fmt.Printf("%s must give: %d;\n", name, imbalance*-1)
		} else if imbalance > 0 {
			fmt.Printf("%s must receive: %d;\n", name, imbalance)
		}
	}
}

func readBills(filename string) ([]Bill, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var bills []Bill
	if err = json.Unmarshal(byteValue, &bills); err != nil {
		return nil, err
	}

	for _, bill := range bills {
		if err = validateBill(bill); err != nil {
			return nil, err
		}
	}

	return bills, nil
}

func validateBill(bill Bill) error {
	sum := 0
	for _, person := range bill.Involved {
		sum += person.Amount
	}
	if sum != 0 {
		return fmt.Errorf("bill %s has a non-zero sum: %d", bill.Name, sum)
	}
	return nil
}

func calculateImbalances(bills []Bill) map[string]int {
	imbalances := make(map[string]int)

	for _, bill := range bills {
		for _, person := range bill.Involved {
			imbalances[person.Name] += person.Amount
		}
	}

	return imbalances
}
