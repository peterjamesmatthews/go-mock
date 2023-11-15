package app

func (application *Application) CoolAlgorithm(x int) int {
	var output int = x

	output = application.Calculator.Add(output, 0)
	output = application.Calculator.Subtract(output, 0)
	output = application.Calculator.Multiply(output, 1)
	output = application.Calculator.Divide(output, 1)

	return output
}
