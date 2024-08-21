package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := prompt("Revenue: ")
	expenses := prompt("Expenses: ")
	taxRate := prompt("Tax rate (e.g. 25.8): ")

	earningsBeforeTax, earningsAfterTax, profitRatio := calculateFinancials(revenue, expenses, taxRate)
	writeOutputToFile(earningsBeforeTax, earningsAfterTax, profitRatio)

	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(profitRatio)
}

func prompt(prompt string) float64 {
	fmt.Print(prompt)

	var input float64
	fmt.Scan(&input)

	if input <= 0 {
		panic("Invalid amount entered.")
	}

	return input
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	profitRatio := ebt / profit

	return ebt, profit, profitRatio
}

func writeOutputToFile(ebt, profit, ratio float64) {
	os.WriteFile("results.txt", []byte(fmt.Sprintf("EBT: %.2f\nProfit: %.2f\n Ratio: %.3f\n", ebt, profit, ratio)), 0644)
}
