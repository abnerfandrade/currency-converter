package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rate struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Please provide two arguments: a number and a currency code.")
		return
	}

	value, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Println("The first argument is not a valid number")
		return
	}
	var brl float64 = value
	var currency string = strings.ToUpper(args[1])

	rawRates, err := os.ReadFile("rates.json")
	if err != nil {
		fmt.Println("Error reading rates file:", err)
		return
	}

	var rates Rate
	err = json.Unmarshal(rawRates, &rates)
	if err != nil {
		fmt.Println("Error unmarshalling rates:", err)
		return
	}

	if _, ok := rates.Rates[currency]; !ok {
		fmt.Println("Currency not found in rates.")
		return
	}

	result := brl * rates.Rates[currency]
	fmt.Printf("%.2f\n", result)
}
