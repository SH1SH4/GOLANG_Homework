package main

import (
	"flag"
	"fmt"
	"os"
	funcs "uniq/uniq"
)

func parseFlags() funcs.Options {
	options := funcs.Options{}

	flag.BoolVar(&options.Count, "c", false, "Подсчитать количество встречаний строки")
	flag.BoolVar(&options.Duplicates, "d", false, "Вывести только повторяющиеся строки")
	flag.BoolVar(&options.Unique, "u", false, "Вывести только уникальные строки")
	flag.BoolVar(&options.IgnoreCase, "i", false, "Игнорировать регистр")
	flag.IntVar(&options.NumFields, "f", 0, "Игнорировать первые N полей")
	flag.IntVar(&options.NumChars, "s", 0, "Игнорировать первые N символов")

	flag.Parse()

	return options
}

func main() {
	options := parseFlags()

	if (options.Count && options.Duplicates) || (options.Count && options.Unique) || (options.Duplicates && options.Unique) {
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

	funcs.Uniq(input, output, options)
}
