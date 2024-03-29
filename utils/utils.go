package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// GetFilename returns first command line argument
func GetFilename() string {
	return os.Args[1]
}

// ReadInput returns the contents of filename as a string.
func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
}

// ExtractInts returns int array from string
func ExtractInts(s string) []int {
	fs := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r) && r != '-'
	})
	ints := make([]int, 0, len(fs))
	for _, w := range fs {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}
