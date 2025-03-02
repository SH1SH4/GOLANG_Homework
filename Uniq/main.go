package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cFlag := flag.Bool("c", false, "Подсчитать количество встречаний строки")
	dFlag := flag.Bool("d", false, "Вывести только повторяющиеся строки")
	uFlag := flag.Bool("u", false, "Вывести только уникальные строки")
	iFlag := flag.Bool("i", false, "Игнорировать регистр")
	fFlag := flag.Int("f", 0, "Игнорировать первые N полей")
	sFlag := flag.Int("s", 0, "Игнорировать первые N символов")
	flag.Parse()

	if (*cFlag && *dFlag) || (*cFlag && *uFlag) || (*dFlag && *uFlag) {
		fmt.Fprintln(os.Stderr, "Ошибка: флаги -c, -d, -u нельзя использовать одновременно")
		flag.Usage()
		os.Exit(1)
	}

	var input *os.File = os.Stdin
	var output *os.File = os.Stdout

	args := flag.Args()
	if len(args) > 0 {
		var err error
		input, err = os.Open(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка открытия входного файла:", err)
			os.Exit(1)
		}
		defer input.Close()
	}

	if len(args) > 1 {
		var err error
		output, err = os.Create(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка создания выходного файла:", err)
			os.Exit(1)
		}
		defer output.Close()
	}

	Choose(input, output, *cFlag, *dFlag, *uFlag, *iFlag, *fFlag, *sFlag)
}
