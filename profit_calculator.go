package main

import "fmt"

func main() {
	revenue := prompt("Revenue: ")
	expenses := prompt("Expenses: ")
	taxRate := prompt("Tax rate (e.g. 25.8): ")

	earningsBeforeTax, earningsAfterTax, profitRatio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(profitRatio)
}

func prompt(prompt string) float64 {
	fmt.Print(prompt)

	var input float64
	fmt.Scan(&input)

	return input
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	profitRatio := ebt / profit

	return ebt, profit, profitRatio
}
