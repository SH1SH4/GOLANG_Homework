package uniq

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func proccessLine(line string, ignoreCase bool, numFields, numChars int) (string, error) {
	if ignoreCase {
		line = strings.ToLower(line)
	}
	// fmt.Println(line)
	if numFields > 0 {
		fields := strings.Fields(line)
		if numFields >= len(fields) {
			return "", nil
		}
		line = strings.Join(fields[numFields:], " ")
	}
	if numChars > len(line) {
		return "", nil
	}
	if numChars > 0 {
		return line[numChars:], nil
	}
	return line, nil
}

func parseLines(input io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(input)
	total := []string{}
	for scanner.Scan() {
		originalLine := scanner.Text()
		total = append(total, originalLine)
	}
	return total, nil
}

func Uniq(input io.Reader, output io.Writer, count, duplicates, unique, ignoreCase bool, numFields, numChars int) error {
	origLines, _ := parseLines(input)
	proccessedLines := []string{}

	for _, v := range origLines {
		temp, err := proccessLine(v, ignoreCase, numFields, numChars)
		if err != nil {
			return err
		}
		proccessedLines = append(proccessedLines, temp)
	}

	writer := bufio.NewWriter(output)
	defer writer.Flush()

	if count {
		counter := 1
		var temp string
		for i := range proccessedLines {
			if i+1 < len(proccessedLines) && proccessedLines[i+1] == proccessedLines[i] {
				if counter == 1 {
					temp = origLines[i]
				}
				counter++
				continue
			}
			if counter == 1 {
				temp = origLines[i]
			}
			fmt.Fprintf(writer, "%v %s\n", counter, temp)
			counter = 1
		}
		return nil
	}

	if duplicates {
		duplicateFlag := false
		var temp string
		for i := range proccessedLines {
			if i+1 < len(proccessedLines) && proccessedLines[i+1] == proccessedLines[i] {
				if !duplicateFlag {
					temp = origLines[i]
				}
				duplicateFlag = true
				continue
			}
			if !duplicateFlag {
				duplicateFlag = true
				continue
			}
			fmt.Fprintf(writer, "%s\n", temp)
			duplicateFlag = false
		}
		return nil
	}
	if unique {
		duplicateFlag := false
		for i := range proccessedLines {
			if i+1 < len(proccessedLines) && proccessedLines[i+1] == proccessedLines[i] {
				duplicateFlag = true
				continue
			}
			if duplicateFlag {
				duplicateFlag = false
				continue
			}
			fmt.Fprintf(writer, "%s\n", origLines[i])
			duplicateFlag = false
		}

		return nil
	}

	for i := range proccessedLines {
		if i != 0 && proccessedLines[i-1] == proccessedLines[i] {
			continue
		}
		fmt.Fprintf(writer, "%s\n", origLines[i])
	}

	return nil
}
