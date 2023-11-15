package calculator

type Calculatorer interface {
	Add(int, int) int
	Subtract(int, int) int
	Multiply(int, int) int
	Divide(int, int) int
}
