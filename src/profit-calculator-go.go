package main

import (
	"fmt"
)

func main() {
	// // initialize variables
	// revenue, err  := strconv.ParseFloat(getUserInput("Revenue: "), 64)
	// if err != nil {
	// 	fmt.Print("Error:", err)
	// }
	// expenses, err := strconv.ParseFloat(getUserInput("Expenses: "), 64)
	// if err != nil {
	// 	fmt.Print("Error:", err)
	// }
	// taxrate, err := strconv.ParseFloat(getUserInput("Tax rate:"), 64)
	// if err != nil {
	// 	fmt.Print("Error:", err)
	// }

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expendses: ")
	taxrate := getUserInput("Tax Rate: ")

	// calculate variables
	ebt, profit, ratio := calculateTaxesProfitRatio(revenue, expenses, taxrate)

	// Outputs information
	outputText(ebt, profit, ratio)
	

}

func outputText(ebt, profit, ratio float64) {
	fmt.Printf(`
	Earnings before tax is: $%.2f

	Profit is: $%.2f
	
	Ratio (EBT/Profit) is: %.2f
	`, ebt, profit, ratio)
}

func calculateTaxesProfitRatio(revenue, expenses, taxrate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	taxes := ebt * (taxrate / 100)
	profit = ebt - taxes
	ratio = ebt / profit
	return  
}

func getUserInput(prompt string) (userInput float64) {
	fmt.Print(prompt)
	fmt.Scan(&userInput)
	return userInput
}
