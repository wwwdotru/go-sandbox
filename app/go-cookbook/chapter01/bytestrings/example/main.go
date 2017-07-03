package main

import (
	"fmt"

	"github.com/wwwdotru/go-sandbox/app/go-cookbook/chapter01/bytestrings"
)

func main() {
	err := bytestrings.WorkWithBuffer()

	if err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println("----------------------------------------------------------")
	bytestrings.SearchString()

	fmt.Println()
	fmt.Println("----------------------------------------------------------")
	bytestrings.ModifyString()

	fmt.Println()
	fmt.Println("----------------------------------------------------------")
	bytestrings.StringReader()

}
