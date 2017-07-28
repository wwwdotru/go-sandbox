package main

import (
	"fmt"

	"github.com/wwwdotru/go-sandbox/app/go-cookbook/chapter01/csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		fmt.Println("There were some errors there")
		fmt.Println(err)
	}

	if err := csvformat.WriteCSVOutput(); err != nil {
		fmt.Println("There were some errors in WriteCSVOutput()")
		fmt.Println(err)
	}

	if err := csvformat.WriteCSVFile(); err != nil {
		fmt.Println("There were some errors in WriteCSVFile()")
		fmt.Println(err)
	}

}
