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
		count        bool
		duplicates   bool
		unique       bool
		ignoreCase   bool
		numFields    int
		numChars     int
		expectsError bool
	}{
		{"Basic Unique", "a\nb\nc\n", "a\nb\nc\n", false, false, false, false, 0, 0, false},
		{"Count Lines", "a\nb\nc\n", "1 a\n1 b\n1 c\n", true, false, false, false, 0, 0, false},
		{"Duplicate Only", "a\na\nb\n", "a\n", false, true, false, false, 0, 0, false},
		{"Unique Only", "a\nb\na\na\n", "a\nb\n", false, false, true, false, 0, 0, false},
		{"Ignore Case", "A\na\nB\nb\n", "A\nB\n", false, false, false, true, 0, 0, false},
		{"Leading Spaces", "  a\n  a\n b\n", "  a\n b\n", false, false, false, false, 0, 0, false},
		{"Trailing Spaces", "a \na \nb \n", "a \nb \n", false, false, false, false, 0, 0, false},
		{"Multiple Duplicates", "x\nx\nx\n", "x\n", false, true, false, false, 0, 0, false},
		{"All Unique", "q\nw\ne\nr\nt\n", "q\nw\ne\nr\nt\n", false, false, false, false, 0, 0, false},
		{"Empty Input", "", "", false, false, false, false, 0, 0, false},
		{"Consecutive Duplicates", "a\na\na\nb\nb\nc\nc\nc\n", "a\nb\nc\n", false, false, false, false, 0, 0, false},
		{"Count Consecutive Duplicates", "a\na\na\nb\nb\nc\nc\nc\n", "3 a\n2 b\n3 c\n", true, false, false, false, 0, 0, false},
		{"Ignore Case with Count", "A\na\nB\nb\nC\nc\n", "2 A\n2 B\n2 C\n", true, false, false, true, 0, 0, false},
		{"Ignore Case with Unique Only", "A\na\nB\nb\nC\n", "C\n", false, false, true, true, 0, 0, false},
		{"Ignore Case with Duplicates Only", "A\na\nB\nb\nC\nc\n", "A\nB\nC\n", false, true, false, true, 0, 0, false},
		{"Num Fields Skip 2", "one two three\none three three\none two two\n", "one two three\none two two\n", false, false, false, false, 2, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile := strings.NewReader(tt.input)
			outputFile := &strings.Builder{}
			err := Uniq(inputFile, outputFile, tt.count, tt.duplicates, tt.unique, tt.ignoreCase, tt.numFields, tt.numChars)
			if tt.expectsError {
				assert.NotNil(t, err, "Ожидалась ошибка, но её нет")
			} else {
				assert.Nil(t, err, "Не должно быть ошибки")
				assert.Equal(t, tt.expected, outputFile.String(), "Выходные данные не совпадают с ожидаемыми")
			}
		})
	}
}
