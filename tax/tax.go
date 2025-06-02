package tax

type Tax struct {
	Name        string
	Rate        float64
	isInclusive bool
}

func CalculateTotal(price float64, taxes []Tax, isInclusive bool) (float64, float64, map[string]float64) {
	var totalExclusivePrice float64
	taxAmounts := make(map[string]float64)
	var totalTax float64

	if isInclusive {
		// Calculate total tax rate
		totalRate := 0.0
		for _, tax := range taxes {
			totalRate += tax.Rate
		}

		// Calculate totalExclusivePrice using inclusive formula
		totalExclusivePrice = price - (totalRate/(100+totalRate))*price

		// Calculate individual tax amounts
		for _, tax := range taxes {
			taxAmount := (tax.Rate / 100) * totalExclusivePrice
			taxAmounts[tax.Name] = taxAmount
			totalTax += taxAmount
		}
	} else {
		// Exclusive tax calculation
		totalExclusivePrice = price
		for _, tax := range taxes {
			taxAmount := (tax.Rate / 100) * totalExclusivePrice
			taxAmounts[tax.Name] = taxAmount
			totalTax += taxAmount
		}
	}

	totalPrice := totalExclusivePrice + totalTax
	return totalExclusivePrice, totalPrice, taxAmounts
}
