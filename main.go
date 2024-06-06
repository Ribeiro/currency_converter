package main

import (
	"fmt"
)

// CurrencyConverter holds the conversion rates relative to a base currency
type CurrencyConverter struct {
	rates map[string]float64
}

// NewCurrencyConverter creates a new CurrencyConverter with predefined rates
func NewCurrencyConverter() *CurrencyConverter {
	return &CurrencyConverter{
		rates: map[string]float64{
			"USD": 1.0, // Base currency
			"EUR": 0.92,
			"JPY": 155.66,
			"GBP": 0.78,
			"INR": 83.46,
			"CAD": 1.37,
			"NZD": 1.61,
			"TRY": 32.26,
		},
	}
}

// Convert converts an amount from one currency to another
func (c *CurrencyConverter) Convert(amount float64, from string, to string) (float64, error) {
	fromRate, fromExists := c.rates[from]
	toRate, toExists := c.rates[to]

	if !fromExists || !toExists {
		return 0, fmt.Errorf("invalid currency code")
	}

	// Convert amount to base currency (USD) and then to the target currency
	amountInBase := amount / fromRate
	convertedAmount := amountInBase * toRate
	return convertedAmount, nil
}

// IsStrongCurrency determines if a currency is strong or weak based on a threshold
func (c *CurrencyConverter) IsStrongCurrency(currency string, threshold float64) (bool, error) {
	rate, exists := c.rates[currency]
	if !exists {
		return false, fmt.Errorf("invalid currency code")
	}

	// Strong if the rate is above the threshold, otherwise weak
	return rate >= threshold, nil
}

func main() {
	converter := NewCurrencyConverter()

	// Example conversion
	amount, err := converter.Convert(100, "TRY", "USD")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("100 TRY is %.2f USD \n", amount)
	}

	// Example check if a currency is strong
	strong, err := converter.IsStrongCurrency("USD", 1.0) // Using 1.0 as threshold for strong currency
	if err != nil {
		fmt.Println("Error:", err)
	} else if strong {
		fmt.Println("USD is a strong currency.")
	} else {
		fmt.Println("USD is a weak currency.")
	}
}
