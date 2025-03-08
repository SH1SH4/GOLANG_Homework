package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expected     int
		expectsError bool
		errorMsg     string
	}{
		{"Addition", "1+1+1", 3, false, ""},
		{"Subtraction", "10-5", 5, false, ""},
		{"Multiplication", "3*3", 9, false, ""},
		{"Division", "10/2", 5, false, ""},
		{"Integer Division", "10/3", 3, false, ""},
		{"Order of Operations", "2+5*(2+10*5)", 262, false, ""},
		{"Division by Zero", "10/0", 0, true, DivisionByZeroException},
		{"Extra Operator", "1++1", 0, true, WrongFormat},
		{"Extra Operator in Brackets", "1+(1++1)", 0, true, WrongFormat},
		{"Extra Closing Bracket", "1+1)", 0, true, ExtraBrackets},
		{"Extra Opening Bracket", "(1+1", 0, true, ExtraBrackets},
		{"Invalid Character", "1+а+1", 0, true, WrongFormat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Calculate(tt.input)
			if tt.expectsError {
				assert.NotNil(t, err, "Ожидалась ошибка, но её нет")
				assert.EqualError(t, err, tt.errorMsg, "Ошибка не соответствует ожидаемой")
				assert.Equal(t, 0, res, "Результат при ошибке должен быть 0")
			} else {
				assert.Nil(t, err, "Не должно быть ошибки")
				assert.Equal(t, tt.expected, res, "Результат не совпадает с ожидаемым")
			}
		})
	}
}
