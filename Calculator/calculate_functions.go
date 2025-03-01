package main

import (
	"fmt"
	"slices"
)

func operation(a, b int, op byte) (int, error) {
	switch op {
	case '+':
		return a + b, nil
	case '-':
		return a - b, nil
	case '*':
		return a * b, nil
	case '/':
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	}
	return 0, fmt.Errorf("неизвестный оператор")
}

func priority(op byte) int {
	if op == '+' || op == '-' {
		return 1
	} else if op == '*' || op == '/' {
		return 2
	}
	return 0
}

func Calculate(line string) (int, error) {
	var values []int
	var operators []byte
	i := 0
	lastWasOperator := false
	for i < len(line) {
		char := line[i]
		if '0' <= char && char <= '9' {
			num := int(char - '0')
			for i+1 < len(line) && '0' <= line[i+1] && line[i+1] <= '9' {
				num = num*10 + int(line[i+1]-'0')
				fmt.Printf("%v\n", num)
				i++
			}
			values = append(values, num)
			lastWasOperator = false
		} else if char == '(' {
			operators = append(operators, line[i])
			lastWasOperator = false
		} else if char == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				a := values[len(values)-1]
				b := values[len(values)-2]
				op := operators[len(operators)-1]
				values = values[:len(values)-2]
				operators = operators[:len(operators)-1]
				temp, err := operation(a, b, op)
				if err != nil {
					return 0, err
				}
				values = append(values, temp)
				lastWasOperator = false
			}
			operators = operators[:len(operators)-1]
		} else if priority(char) != 0 {
			if lastWasOperator {
				return 0, fmt.Errorf("неверный формат: два оператора подряд")
			}
			for len(operators) != 0 && priority(char) <= priority(operators[len(operators)-1]) {
				a := values[len(values)-1]
				b := values[len(values)-2]
				op := operators[len(operators)-1]
				values = values[:len(values)-2]
				operators = operators[:len(operators)-1]
				temp, err := operation(a, b, op)
				if err != nil {
					return 0, err
				}
				values = append(values, temp)
			}
			operators = append(operators, char)
			lastWasOperator = true
		} else {
			return 0, fmt.Errorf("неизвестный символ '%c'", char)
		}
		i++
	}
	if len(values)-1 != len(operators) {
		return 0, fmt.Errorf("%v : %v", len(values), len(operators))
		// return 0, fmt.Errorf("неверный формат")
	}
	if slices.Contains(operators, '(') {
		return 0, fmt.Errorf("незакрытая скобка")
	}
	for len(operators) > 0 {
		b := values[len(values)-1]
		a := values[len(values)-2]
		op := operators[len(operators)-1]

		values = values[:len(values)-2]
		operators = operators[:len(operators)-1]
		temp, err := operation(a, b, op)
		if err != nil {
			return 0, err
		}
		values = append(values, temp)
	}
	return values[0], nil
}
