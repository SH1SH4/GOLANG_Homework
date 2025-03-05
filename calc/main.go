package main

import (
	calculator "calc/Calculator"
	"fmt"
	// "os"
)

func main() {
	var line string
	fmt.Print("Введите пример без пробелов :")
	// fmt.Fscan(os.Stdin, &line)
	line = "2+5*(2+10*5)"
	answer, err := calculator.Calculate(line)
	if err == nil {
		fmt.Printf("%v=%v", line, answer)
	}
	if err != nil {
		fmt.Printf("Ошибка: %v", err)
	}
}
