package csvformat

import (
	"encoding/csv"
	"io"
	"os"
)

// Book is a good thing
type Book struct {
	Author string
	Title  string
}

// Books is a collection of objects Book
type Books []Book

// ToCSV takes a set of books and writes it to the io.Writer
func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}

	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}
	n.Flush()

	return n.Error()
}

// WriteCSVOutput example of using a function and write it to std.Out
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "This is the first book",
			Title:  "Some gal",
		},
		Book{
			Author: "This is the second book",
			Title:  "Some guy",
		},
	}
	return b.ToCSV(os.Stdout)
}

// WriteCSVFile example of using a function and write it to file
func WriteCSVFile() error {
	b := Books{
		Book{
			Author: "This is the first book",
			Title:  "Some gal",
		},
		Book{
			Author: "This is the second book",
			Title:  "Some guy",
		},
	}

	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	return b.ToCSV(f)
}
