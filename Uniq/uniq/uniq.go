package uniq

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Options struct {
	Count      bool
	Duplicates bool
	Unique     bool
	IgnoreCase bool
	NumFields  int
	NumChars   int
}

func proccessLine(line string, ignoreCase bool, numFields, numChars int) string {
	if ignoreCase {
		line = strings.ToLower(line)
	}
	if numFields > 0 {
		fields := strings.Fields(line)
		if numFields >= len(fields) {
			return ""
		}
		line = strings.Join(fields[numFields:], " ")
	}
	if numChars > len(line) {
		return ""
	}
	if numChars > 0 {
		return line[numChars:]
	}
	return line
}

func parseLines(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	total := []string{}
	for scanner.Scan() {
		originalLine := scanner.Text()
		total = append(total, originalLine)
	}
	return total
}

func Uniq(input io.Reader, output io.Writer, options Options) {
	origLines := parseLines(input)
	proccessedLines := []string{}

	for _, v := range origLines {
		temp := proccessLine(v, options.IgnoreCase, options.NumFields, options.NumChars)
		proccessedLines = append(proccessedLines, temp)
	}

	writer := bufio.NewWriter(output)
	defer writer.Flush()

	counter := 1
	var temp string
	for i := range proccessedLines {
		if counter == 1 {
			temp = origLines[i]
		}
		if i+1 < len(proccessedLines) && proccessedLines[i+1] == proccessedLines[i] {
			counter++
			continue
		}
		if !(options.Count || options.Unique || options.Duplicates) {
			fmt.Fprintf(writer, "%s\n", temp)
		}
		if options.Duplicates && counter > 1 {
			fmt.Fprintf(writer, "%s\n", temp)
		}
		if options.Unique && counter == 1 {
			fmt.Fprintf(writer, "%s\n", temp)
		}
		if options.Count {
			fmt.Fprintf(writer, "%v %s\n", counter, temp)
		}
		counter = 1
	}
}
