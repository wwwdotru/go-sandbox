package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// SearchString shows few methods for searching string
func SearchString() {
	s := "this is a test"

	fmt.Println(strings.Contains(s, "this"))

	fmt.Println(strings.ContainsAny(s, "bca"))

	fmt.Println(strings.HasPrefix(s, "this"))

	fmt.Println(strings.HasSuffix(s, "test"))
}

// ModifyString modifies string in few ways
func ModifyString() {
	s := "simple string"

	fmt.Println(strings.Split(s, " "))

	fmt.Println(strings.Title(s))

	s = " simple string "
	fmt.Println(strings.TrimSpace(s))
}

// StringReader demonstrates how to create an io.Reader interface quickly with a string
func StringReader() {
	s := "simple string"
	r := strings.NewReader(s)

	io.Copy(os.Stdout, r)
}
