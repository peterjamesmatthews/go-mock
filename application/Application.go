package application

import "pjm.dev/mock/calculator"

// Application holds the main application state and functionality.
type Application struct {
	// Calculator is the application's interface with the calculator service.
	Calculator calculator.Calculatorer
}
