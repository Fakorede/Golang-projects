package simplemath

// Sum : sums up slice values provided as argument.
func Sum(values ...float64) float64 {
	total := 0.0

	for _, value := range values {
		total += value
	}

	return total
}

// Multiply : multiplies argument p1 and argument p2
func Multiply(p1, p2 float64) float64 {
	return p1 * p2
}

// Divide : divides argument p1 by argument p2
func Divide(p1, p2 float64) float64 {
	if p2 == 0 {
		return 0
	}

	return p1 / p2
}

// Subtract : subtracts the provided arguments
func Subtract(p1, p2 float64) float64 {
	return p1 - p2
}

// Add : adds the provided arguments
func Add(p1, p2 float64) float64 {
	return p1 + p2
}
