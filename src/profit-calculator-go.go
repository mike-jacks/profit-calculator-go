package main

import (
	"fmt"
)

func main() {
	// initialize variables
	var revenue float64
	var expenses float64
	var taxrate float64

	// get values for variables above from user
	fmt.Print("revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("tax rate: ")
	fmt.Scan(&taxrate)

	// calculate variables
	earningsbeforetax := revenue - expenses
	taxes := earningsbeforetax * (taxrate / 100)
	profit := earningsbeforetax - taxes
	ratio := earningsbeforetax / profit

	fmt.Println("")
	fmt.Println("Earnings before tax is: $",earningsbeforetax)
	fmt.Println("Profit is: $",profit)
	fmt.Println("Ratio (EBT/Profit) is:",ratio)

}



