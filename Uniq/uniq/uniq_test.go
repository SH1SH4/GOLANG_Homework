package uniq

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expected     string
		options      Options
		expectsError bool
	}{
		{"Basic Unique", "a\nb\nc\n", "a\nb\nc\n", Options{}, false},
		{"Count Lines", "a\nb\nc\n", "1 a\n1 b\n1 c\n", Options{Count: true}, false},
		{"Duplicate Only", "a\na\nb\n", "a\n", Options{Duplicates: true}, false},
		{"Unique Only", "a\nb\na\na\n", "a\nb\n", Options{Unique: true}, false},
		{"Ignore Case", "A\na\nB\nb\n", "A\nB\n", Options{IgnoreCase: true}, false},
		{"Leading Spaces", "  a\n  a\n b\n", "  a\n b\n", Options{}, false},
		{"Trailing Spaces", "a \na \nb \n", "a \nb \n", Options{}, false},
		{"Multiple Duplicates", "x\nx\nx\n", "x\n", Options{Duplicates: true}, false},
		{"All Unique", "q\nw\ne\nr\nt\n", "q\nw\ne\nr\nt\n", Options{}, false},
		{"Empty Input", "", "", Options{}, false},
		{"Consecutive Duplicates", "a\na\na\nb\nb\nc\nc\nc\n", "a\nb\nc\n", Options{}, false},
		{"Count Consecutive Duplicates", "a\na\na\nb\nb\nc\nc\nc\n", "3 a\n2 b\n3 c\n", Options{Count: true}, false},
		{"Ignore Case with Count", "A\na\nB\nb\nC\nc\n", "2 A\n2 B\n2 C\n", Options{Count: true, IgnoreCase: true}, false},
		{"Ignore Case with Unique Only", "A\na\nB\nb\nC\n", "C\n", Options{Unique: true, IgnoreCase: true}, false},
		{"Ignore Case with Duplicates Only", "A\na\nB\nb\nC\nc\n", "A\nB\nC\n", Options{Duplicates: true, IgnoreCase: true}, false},
		{"Num Fields Skip 2", "one two three\none three three\none two two\n", "one two three\none two two\n", Options{NumFields: 2}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile := strings.NewReader(tt.input)
			outputFile := &strings.Builder{}
			Uniq(inputFile, outputFile, tt.options)
			assert.Equal(t, tt.expected, outputFile.String(), "Выходные данные не совпадают с ожидаемыми")
		})
	}
}
