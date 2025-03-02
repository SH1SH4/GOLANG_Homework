package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countLine(line []string, item string) int {
	counter := 0
	for _, i := range line {
		if i == item {
			counter++
		}
	}
	return counter
}

func proccessLine(line string, ignoreCase bool, numFields, numChars int) (string, error) {
	if ignoreCase {
		line = strings.ToLower(line)
	}
	if numFields > 0 {
		fields := strings.Fields(line)
		if numFields > len(fields) {
			return "", fmt.Errorf("полей для пропуска больше, чем полей в строке всего")
		}
		line = strings.Join(fields[numFields:], " ")
	}
	if numChars > 0 {
		line = line[numChars:]
	}
	return line, nil
}

func uniqStrings(input *os.File, output *os.File, ignoreCase bool, numFields, numChars int) {

	var uniqString []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		processedLine, _ := proccessLine(scanner.Text(), ignoreCase, numFields, numChars)
		if len(uniqString) == 0 || uniqString[len(uniqString)-1] != processedLine {
			uniqString = append(uniqString, processedLine)
		}
	}

	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for _, line := range uniqString {
		fmt.Fprintf(writer, "%s\n", line)
	}
}

func countUniqStrings(input *os.File, output *os.File, ignoreCase bool, numFields, numChars int) {
	scanner := bufio.NewScanner(input)
	var uniqStringCounter []int
	var uniqString []string
	for scanner.Scan() {
		processedLine, _ := proccessLine(scanner.Text(), ignoreCase, numFields, numChars)
		if len(uniqString) == 0 || uniqString[len(uniqString)-1] != processedLine {
			uniqString = append(uniqString, processedLine)
			uniqStringCounter = append(uniqStringCounter, 1)
		} else {
			uniqStringCounter[len(uniqStringCounter)-1]++
		}
	}

	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for i, line := range uniqString {
		fmt.Fprintln(writer, uniqStringCounter[i], line)
	}
}

func duplicateStrings(input *os.File, output *os.File, ignoreCase bool, numFields, numChars int) {
	var strings []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		processedLine, _ := proccessLine(scanner.Text(), ignoreCase, numFields, numChars)
		strings = append(strings, processedLine)
	}
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for _, line := range strings {
		if countLine(strings, line) > 1 {
			fmt.Fprintf(writer, "%s\n", line)
		}
	}
}

func onlyUniqStrings(input *os.File, output *os.File, ignoreCase bool, numFields, numChars int) {
	var strings []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		processedLine, _ := proccessLine(scanner.Text(), ignoreCase, numFields, numChars)
		strings = append(strings, processedLine)
	}
	writer := bufio.NewWriter(output)
	defer writer.Flush()

	for _, line := range strings {
		if countLine(strings, line) == 1 {
			fmt.Fprintf(writer, "%s\n", line)
		}
	}
}

func Choose(input *os.File, output *os.File, count, duplicates, unique, ignoreCase bool, numFields, numChars int) {
	if count {
		countUniqStrings(input, output, ignoreCase, numFields, numChars)
	} else if duplicates {
		duplicateStrings(input, output, ignoreCase, numFields, numChars)
	} else if unique {
		onlyUniqStrings(input, output, ignoreCase, numFields, numChars)
	} else {
		uniqStrings(input, output, ignoreCase, numFields, numChars)
	}
}
