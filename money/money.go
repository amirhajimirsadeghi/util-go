package money

import (
	"fmt"
	"strconv"
	"strings"
)

type Money struct {
	cents int
}

func FromCents(cents int) Money {
	return Money{cents: cents}
}

func FromString(s string) (Money, error) {
	var money Money

	// Remove dollar sign and commas
	dollars := s
	dollars = strings.ReplaceAll(dollars, "$", "")
	dollars = strings.ReplaceAll(dollars, ",", "")

	// Determine if the amount is negative and clean up
	var isNegative bool
	if strings.HasPrefix(dollars, "-") {
		isNegative = true
		dollars = dollars[1:]
	}

	// Split dollars and cents
	parts := strings.Split(dollars, ".")
	if len(parts) > 2 {
		return money, fmt.Errorf("invalid format")
	}

	// Calculate cents
	centsPart := "0"
	if len(parts) == 2 && parts[1] != "" {
		centsPart = parts[1]
		if len(centsPart) == 1 {
			centsPart += "0" // For single-digit cents
		}
	}
	if len(centsPart) > 2 {
		return money, fmt.Errorf("too many digits after decimal")
	}

	// Calculate dollars
	dollarsPart := parts[0]
	if dollarsPart == "" {
		dollarsPart = "0"
	}

	// Convert dollars and cents to integers
	dollarsInt, err := strconv.Atoi(dollarsPart)
	if err != nil {
		return money, err
	}
	centsInt, err := strconv.Atoi(centsPart)
	if err != nil {
		return money, err
	}

	// Combine dollars and cents and return
	money.cents = dollarsInt*100 + centsInt
	if isNegative {
		money.cents = -money.cents
	}
	return money, nil
}

func (m Money) Cents() int {
	return m.cents
}

func (m Money) String() string {
	// Generate dollar string
	dollars := m.cents / 100
	if dollars < 0 {
		dollars = -dollars
	}
	dollarStr := fmt.Sprintf("%d", dollars)
	i := -1
	for dollars >= 1000 {
		i += 4
		dollarStr = dollarStr[:len(dollarStr)-i] + "," + dollarStr[len(dollarStr)-i:]
		dollars /= 1000
	}

	// Generate cent string
	cents := m.cents % 100
	if cents < 0 {
		cents = -cents
	}
	centStr := fmt.Sprintf("%02d", cents)

	// Combine dollar and cent strings and attach symbols and return
	s := "$" + dollarStr + "." + centStr
	if m.cents < 0 {
		s = "-" + s
	}

	return s
}
