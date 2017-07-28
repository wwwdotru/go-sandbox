package csvformat

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Movie will hold parced csv
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV is an example of reading csv file into slice
func ReadCSV(b io.Reader) ([]Movie, error) {
	r := csv.NewReader(b)

	// r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// Header
	// h, err := r.Read()
	// fmt.Printf("%v", h)
	// if err != nil && err != io.EOF {
	// 	return nil, err
	// }

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}
		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}

	return movies, nil
}

// AddMoviesFromText uses the CSV parser with the string
func AddMoviesFromText() error {
	f, err := os.Open("test.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	m, err := ReadCSV(f)
	if err != nil {
		return err
	}
	for _, movie := range m {
		fmt.Printf("%#v\n", movie)
	}

	in := `
- headers
moo,doo,102
roo,goo,30
`
	fmt.Println()
	fmt.Println("Now form the string buffer")
	fmt.Println()

	b := bytes.NewBufferString(in)
	m, err = ReadCSV(b)
	if err != nil {
		return err
	}
	for _, movie := range m {
		fmt.Printf("%#v\n", movie)
	}

	return nil
}
