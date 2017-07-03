package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkWithBuffer will make use of the buffer created by the Buffer function
func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"

	// convert Buffer back to string
	b := Buffer(rawString)
	fmt.Println(b.String())

	// we can use generic reader functions( we created it in buffer.go)
	s, err := toString(b)
	if err != nil {
		return err
	}
	fmt.Println(s)
	// we can also take our bytes and create a bytes reader
	// these readers implement io.Reader, io.ReaderAt,
	// io.WriterTo, io.Seeker, io.ByteScanner, and
	// io.RuneScanner interfaces
	reader := bytes.NewReader([]byte(rawString))

	// we can also plug it into a scanner that allows
	// buffered reading and tokenzation
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	// iterate over all of the scan events
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}

	return nil
}
