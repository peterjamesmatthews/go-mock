package app_test

import (
	"testing"

	"go.uber.org/mock/gomock"
	"pjm.dev/mock/app"
	"pjm.dev/mock/mock_calculator"
)

func TestCoolAlgorithm(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCalculator := mock_calculator.NewMockCalculatorer(ctrl)

	number := 100

	mockCalculator.
		EXPECT().
		Add(number, 0).
		Times(1).
		Return(number)

	mockCalculator.
		EXPECT().
		Subtract(number, 0).
		Times(1).
		Return(number)

	mockCalculator.
		EXPECT().
		Multiply(number, 1).
		Times(1).
		Return(number)

	mockCalculator.
		EXPECT().
		Divide(number, 1).
		Times(1).
		Return(number)

	application := app.Application{Calculator: mockCalculator}
	if application.CoolAlgorithm(number) != number {
		t.Fail()
	}
}
