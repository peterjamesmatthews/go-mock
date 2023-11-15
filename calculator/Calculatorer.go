package calculator

// Interface that the external calculator service implements.
type Calculatorer interface {
	// Add returns the sum of two ints.
	Add(int, int) int
	// Subtract returns the difference of two ints.
	Subtract(int, int) int
	// Multiply returns the product of two ints.
	Multiply(int, int) int
	// Divide returns the quotient of two ints.
	Divide(int, int) int
}
