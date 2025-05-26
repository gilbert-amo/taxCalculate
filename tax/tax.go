package tax

type Tax struct {
	Name string
	Rate float64
}

func CalculateTotal(price float64, taxes []Tax, isInclusive bool) (float64, float64, map[string]float64) {
	var subtotal float64
	taxAmounts := make(map[string]float64)
	var totalTax float64

	if isInclusive {
		// Calculate total tax rate
		totalRate := 0.0
		for _, tax := range taxes {
			totalRate += tax.Rate
		}

		// Calculate subtotal using inclusive formula
		subtotal = price - (totalRate/(100+totalRate))*price

		// Calculate individual tax amounts
		for _, tax := range taxes {
			taxAmount := (tax.Rate / 100) * subtotal
			taxAmounts[tax.Name] = taxAmount
			totalTax += taxAmount
		}
	} else {
		// Exclusive tax calculation
		subtotal = price
		for _, tax := range taxes {
			taxAmount := (tax.Rate / 100) * subtotal
			taxAmounts[tax.Name] = taxAmount
			totalTax += taxAmount
		}
	}

	totalPrice := subtotal + totalTax
	return subtotal, totalPrice, taxAmounts
}
