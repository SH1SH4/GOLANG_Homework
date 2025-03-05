package Test

import (
	calculator "calc/Calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	assert := assert.New(t)
	res, err := calculator.Calculate("1+1+1")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(3, res, "1+1+1 должно быть равно 3")
}
func TestMinus(t *testing.T) {
	assert := assert.New(t)
	// Тест на вычитание
	res, err := calculator.Calculate("10-5")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(5, res, "10-5 должно быть равно 5")
}
func TestMultiply(t *testing.T) {
	assert := assert.New(t)
	// Тест на умножение
	res, err := calculator.Calculate("3*3")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(9, res, "3*3 должно быть равно 9")
}
func TestDivision(t *testing.T) {
	assert := assert.New(t)
	// Тест на деление
	res, err := calculator.Calculate("10/2")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(5, res, "10/2 должно быть равно 5")
}
func TestDivisionRemainder(t *testing.T) {
	assert := assert.New(t)
	// Тест на деление с остатком (целочисленное деление)
	res, err := calculator.Calculate("10/3")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(3, res, "10/3 должно быть равно 3")
}

func TestOrderOfOperations(t *testing.T) {
	assert := assert.New(t)
	// Тест на порядок операций
	res, err := calculator.Calculate("2+5*(2+10*5)")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(262, res, "2+5*(2+10*5) = 262")
}
func TestDivisionByZero(t *testing.T) {
	assert := assert.New(t)
	// Тест на ошибку: деление на ноль
	res, err := calculator.Calculate("10/0")
	assert.NotNil(err, "Должна быть ошибка")
	assert.EqualError(err, calculator.DivisionByZeroException, "Не та ошибка")
	assert.Equal(0, res, "Результат при делении на ноль должен быть равен 0")
}

func TestIncorrectException(t *testing.T) {
	assert := assert.New(t)
	// Тест на лишний оператор
	res, err := calculator.Calculate("1++1")
	assert.NotNil(err, "Должна быть ошибка")
	assert.EqualError(err, calculator.WrongFormat, "Не та ошибка")
	assert.Equal(0, res, "Результат при некорректном выражении должен быть равен 0")
}

func TestIncorrectException2(t *testing.T) {
	assert := assert.New(t)
	// Тест на лишний оператор внутри скобок
	res, err := calculator.Calculate("1+(1++1)")
	assert.NotNil(err, "Должна быть ошибка")
	assert.EqualError(err, calculator.WrongFormat, "Не та ошибка")
	assert.Equal(0, res, "Результат при некорректном выражении должен быть равен 0")
}

func TestIncorrectException3(t *testing.T) {
	assert := assert.New(t)
	// Тест на лишние скобки
	res, err := calculator.Calculate("1+1)")
	assert.NotNil(err, "Должна быть ошибка")
	assert.EqualError(err, calculator.ExtraBrackets, "Не та ошибка")
	assert.Equal(0, res, "Результат при некорректном выражении должен быть равен 0")
}

func TestIncorrectException4(t *testing.T) {
	assert := assert.New(t)
	// Тест на лишние скобки
	res, err := calculator.Calculate("(1+1")
	assert.NotNil(err, "Должна быть ошибка")
	assert.EqualError(err, calculator.ExtraBrackets, "Не та ошибка")
	assert.Equal(0, res, "Результат при некорректном выражении должен быть равен 0")
}

func TestWrongSymbols(t *testing.T) {
	assert := assert.New(t)
	// Тест на сторонние символы в выражении
	res, err := calculator.Calculate("1+а+1")
	assert.NotNil(err, "Должна быть ошибка")
	assert.EqualError(err, calculator.WrongFormat, "Не та ошибка")
	assert.Equal(0, res, "Результат при некорректном выражении должен быть равен 0")
}
