package main

import (
	"fmt"
)

func operation(a, b int, op byte) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}

func priority(op byte) int {
	if op == '+' || op == '-' {
		return 1
	} else if op == '*' || op == '/' {
		return 2
	}
	return 0
}

func calculate(line string) int {
	var values []int
	var operators []byte
	i := 0
	for i < len(line) {
		char := line[i]

		// fmt.Printf("%v : %c\n", line[i], line[i])
		// if char == ' ' {
		// 	i++
		// 	continue
		// }
		if '0' <= char && char <= '9' {
			num := int(char - '0')
			for i+1 < len(line) && '0' <= line[i+1] && line[i+1] <= '9' {
				num = num*10 + int(line[i]-'0')
				i++
			}
			values = append(values, num)
		} else if char == '(' {
			operators = append(operators, line[i])
		} else if char == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				a := values[len(values)-1]
				b := values[len(values)-2]
				fmt.Println(int(a))
				fmt.Println(b)
				op := operators[len(operators)-1]
				values = values[:len(values)-2]
				operators = operators[:len(operators)-1]
				temp := operation(a, b, op)
				fmt.Println(temp)
				values = append(values, temp)
			}
			operators = operators[:len(operators)-1]
		} else if priority(char) != 0 {
			for len(operators) != 0 && priority(char) <= priority(operators[len(operators)-1]) {
				a := values[len(values)-1]
				b := values[len(values)-2]
				fmt.Println(int(a))
				fmt.Println(b)
				op := operators[len(operators)-1]
				values = values[:len(values)-2]
				operators = operators[:len(operators)-1]
				temp := operation(a, b, op)
				fmt.Println(temp)
				values = append(values, temp)
			}
			operators = append(operators, char)
		}
		i++
	}
	for len(operators) > 0 {
		b := values[len(values)-1]
		a := values[len(values)-2]
		op := operators[len(operators)-1]

		values = values[:len(values)-2]
		operators = operators[:len(operators)-1]

		values = append(values, operation(a, b, op))
	}
	return values[0]
}

func main() {
	var line string = "1+2*(2*3+1)"
	// fmt.Fscan(os.Stdin, &line)
	fmt.Printf("Ответ: %v", calculate(line))
}
