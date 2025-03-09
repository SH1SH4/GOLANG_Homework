package main

import (
	funcs "uniq/Funcs"

	"flag"
	"fmt"
	"os"
)

func main() {
	countFlag := flag.Bool("c", false, "Подсчитать количество встречаний строки")
	duplicateFlag := flag.Bool("d", false, "Вывести только повторяющиеся строки")
	uniqFlag := flag.Bool("u", false, "Вывести только уникальные строки")
	ignoreCaseFlag := flag.Bool("i", false, "Игнорировать регистр")
	fieldFlag := flag.Int("f", 0, "Игнорировать первые N полей")
	CharFlag := flag.Int("s", 0, "Игнорировать первые N символов")
	flag.Parse()

	if (*countFlag && *duplicateFlag) || (*countFlag && *uniqFlag) || (*duplicateFlag && *uniqFlag) {
		fmt.Fprintln(os.Stderr, "Ошибка: флаги -c, -d, -u нельзя использовать одновременно")
		flag.Usage()
		return
	}

	var input *os.File = os.Stdin
	var output *os.File = os.Stdout

	args := flag.Args()
	if len(args) > 0 {
		var err error
		input, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка открытия входного файла:", err)
			return
		}
		defer input.Close()
	}

	if len(args) > 1 {
		var err error
		output, err = os.Create(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка создания выходного файла:", err)
			return
		}
		defer output.Close()
	}
	err := funcs.Uniq(input, output, *countFlag, *duplicateFlag, *uniqFlag, *ignoreCaseFlag, *fieldFlag, *CharFlag)
	if err != nil {
		fmt.Println(err)
	}
}
