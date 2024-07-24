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
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
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
			fmt.Printf("%s must give: %.2f;\n", name, imbalance*-1)
		} else if imbalance > 0 {
			fmt.Printf("%s must receive: %.2f;\n", name, imbalance)
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
	sum := 0.0
	numberOfSplitters := 0
	for _, person := range bill.Involved {
		sum += person.Amount
		if person.Amount == 0 {
			numberOfSplitters += 1
		}
	}
	if (sum != 0 && numberOfSplitters == 0) || (numberOfSplitters > 0 && sum <= 0) {
		return fmt.Errorf("bill %s is not valid", bill.Name)
	}
	return nil
}

func calculateImbalances(bills []Bill) map[string]float64 {
	imbalances := make(map[string]float64)

	for _, bill := range bills {
		var splitters []string
		var payedAmount, exactAmount float64
		for _, person := range bill.Involved {
			if person.Amount < 0 {
				exactAmount += person.Amount
				imbalances[person.Name] += person.Amount
			} else if person.Amount > 0 {
				payedAmount += person.Amount
				imbalances[person.Name] += person.Amount
			} else {
				splitters = append(splitters, person.Name)
			}
		}
		if len(splitters) > 0 {
			amount := (payedAmount + exactAmount) / float64(len(splitters))
			for _, splitter := range splitters {
				imbalances[splitter] -= amount
			}
		}
	}

	return imbalances
}
