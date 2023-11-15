package application

// CoolAlgorithm is an important bit of Application functionality that relies on the Calculatorer interface.
func (application *Application) CoolAlgorithm(input int) int {
	var output int = input

	output = application.Calculator.Add(output, 0)
	output = application.Calculator.Subtract(output, 0)
	output = application.Calculator.Multiply(output, 1)
	output = application.Calculator.Divide(output, 1)

	return output
}
