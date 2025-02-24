package main

import (
	"fmt"
	"os"
)

func main() {
	var line string
	fmt.Fscan(os.Stdin, &line)
	for _, char := range line {
		fmt.Printf("%v : %c\n", char, char)
		if '0' <= char && char <= '9' {
			fmt.Println("цифра")
		}

	}
}
