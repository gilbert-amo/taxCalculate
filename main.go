package main

import (
	"bufio"
	"fmt"
	"github.com/gilbert-amo/taxCalculate/tax"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get price input
	fmt.Print("Enter the price: ")
	priceInput, _ := reader.ReadString('\n')
	price, err := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)
	if err != nil {
		fmt.Println("Invalid price entered")
		return
	}

	// Get tax type input
	var taxes []tax.Tax
	for {
		fmt.Print("Enter tax name (or 'done' to finish): ")
		nameInput, _ := reader.ReadString('\n')
		name := strings.TrimSpace(nameInput)
		if strings.ToLower(name) == "done" {
			break
		}

		fmt.Print("Enter tax rate (%): ")
		rateInput, _ := reader.ReadString('\n')
		rate, err := strconv.ParseFloat(strings.TrimSpace(rateInput), 64)
		if err != nil {
			fmt.Println("Invalid rate entered. Please try again.")
			continue
		}

		taxes = append(taxes, tax.Tax{Name: name, Rate: rate})
	}

	if len(taxes) == 0 {
		fmt.Println("No taxes entered. Exiting.")
		return
	}

	// Get calculation type
	fmt.Print("Is tax inclusive? (y/n): ")
	inclusiveInput, _ := reader.ReadString('\n')
	isInclusive := strings.ToLower(strings.TrimSpace(inclusiveInput)) == "y"

	// Calculate and display results
	subtotal, total, taxAmounts := tax.CalculateTotal(price, taxes, isInclusive)

	fmt.Println("\n=== Calculation Results ===")
	fmt.Printf("Original Price: GHS%.2f\n", price)
	fmt.Printf("Subtotal:GHS%.2f\n", subtotal)

	fmt.Println("\nTax Breakdown:")
	for name, amount := range taxAmounts {
		for _, tax := range taxes {
			if tax.Name == name {
				fmt.Printf("- %s (%.2f%%): GHS%.2f\n", name, tax.Rate, amount)
				break
			}
		}
	}

	fmt.Printf("\nTotal: GHS%.2f\n", total)
}
