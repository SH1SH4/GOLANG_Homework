package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	assert := assert.New(t)

	// Тест на сложение
	res, err := Calculate("1+1+1")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(3, res, "1+1+1 должно быть равно 3")

	// Тест на вычитание
	res, err = Calculate("10-5")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(5, res, "10-5 должно быть равно 5")

	// Тест на умножение
	res, err = Calculate("3*3")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(9, res, "3*3 должно быть равно 9")

	// Тест на деление
	res, err = Calculate("10/2")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(5, res, "10/2 должно быть равно 5")

	// Тест на деление с остатком (целочисленное деление)
	res, err = Calculate("10/3")
	assert.Nil(err, "Не должно быть ошибки")
	assert.Equal(3, res, "10/3 должно быть равно 3")

	// Тест на ошибку: деление на ноль
	res, err = Calculate("10/0")
	assert.NotNil(err, "Ошибка при делении на ноль")
	assert.Equal(0, res, "Результат при делении на ноль должен быть равен 0")

	// Тест на неправильное выражение
	res, err = Calculate("1++1")
	assert.NotNil(err, "Ошибка при некорректном выражении")
	assert.Equal(0, res, "Результат при некорректном выражении должен быть равен 0")

	// Тест на сторонние символы в выражении
	res, err = Calculate("1+а+1")
	assert.NotNil(err, "Ошибка при лишних символах")
	// assert.Equal(0, res, "Результат при некорректном выражении должен быть равен 0")
}
