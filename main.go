package main

import (
	"fmt"
	"os"
)

func main() {
	var line string
	fmt.Print("Введите пример без пробелов :")
	fmt.Fscan(os.Stdin, &line)
	// line = "100-5"
	answer, err := Calculate(line)
	if err == nil {
		fmt.Printf("%v=%v", line, answer)
	} else {
		fmt.Printf("Ошибка: %v", err)
	}
}
