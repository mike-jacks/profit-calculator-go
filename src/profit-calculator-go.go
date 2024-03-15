package main

import (
	"fmt"
)

func main() {
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
	_, printError := fmt.Print(prompt)
	if printError != nil {
		print("Error printing")
	}
	_, userInputError := fmt.Scan(&userInput)
	if userInputError != nil {
		print("Error with user input not being a float")
	}
	
	return userInput
}
