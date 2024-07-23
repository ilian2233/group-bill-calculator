package main

import (
	"os"
	"reflect"
	"testing"
)

func TestReadBillsValidFile(t *testing.T) {
	content := `[{"name":"Test Bill","involved":[{"name":"John","amount":10},{"name":"Jane","amount":-10}]}]`
	file, err := os.CreateTemp("", "bills*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	bills, err := readBills(file.Name())
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if len(bills) != 1 || bills[0].Name != "Test Bill" {
		t.Errorf("Unexpected bills content: %+v", bills)
	}
}

func TestReadBillsInvalidFile(t *testing.T) {
	content := `invalid JSON`
	file, err := os.CreateTemp("", "bills*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	_, err = readBills(file.Name())
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestValidateBill(t *testing.T) {
	tests := []struct {
		name     string
		bill     Bill
		hasError bool
	}{
		{
			name: "Zero Sum Bill",
			bill: Bill{
				Name: "Zero Sum Bill",
				Involved: []Person{
					{Name: "John", Amount: -10},
					{Name: "Jane", Amount: 10},
				},
			},
			hasError: false,
		},
		{
			name: "Non-Zero Sum Bill",
			bill: Bill{
				Name: "Non-Zero Sum Bill",
				Involved: []Person{
					{Name: "John", Amount: 10},
					{Name: "Jane", Amount: -5},
				},
			},
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateBill(tt.bill)
			if (err != nil) != tt.hasError {
				t.Errorf("validateBill() error = %v, wantErr %v", err, tt.hasError)
			}
		})
	}
}

func TestCalculateImbalances(t *testing.T) {
	tests := []struct {
		name               string
		bills              []Bill
		expectedImbalances map[string]int
	}{
		{
			name: "Single Bill Multiple People",
			bills: []Bill{
				{
					Name: "Lunch",
					Involved: []Person{
						{Name: "Alice", Amount: 10},
						{Name: "Bob", Amount: -10},
					},
				},
			},
			expectedImbalances: map[string]int{
				"Alice": 10,
				"Bob":   -10,
			},
		},
		{
			name: "Multiple Bills",
			bills: []Bill{
				{
					Name: "Lunch",
					Involved: []Person{
						{Name: "Alice", Amount: 10},
						{Name: "Bob", Amount: -5},
					},
				},
				{
					Name: "Dinner",
					Involved: []Person{
						{Name: "Alice", Amount: -5},
						{Name: "Bob", Amount: 5},
					},
				},
			},
			expectedImbalances: map[string]int{
				"Alice": 5,
				"Bob":   0,
			},
		},
		{
			name:               "No Bills",
			bills:              []Bill{},
			expectedImbalances: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imbalances := calculateImbalances(tt.bills)
			if !reflect.DeepEqual(imbalances, tt.expectedImbalances) {
				t.Errorf("calculateImbalances() got = %v, want %v", imbalances, tt.expectedImbalances)
			}
		})
	}
}
