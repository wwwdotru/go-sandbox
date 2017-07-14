package filedirs

import (
	"bytes"
	"io"
	"os"
	"strings"
)

// Capitalizer example of writing one file into another after reading it
func Capitalizer(f1 *os.File, f2 *os.File) error {
	if _, err := f1.Seek(0, 0); err != nil {
		return err
	}

	var tmp = new(bytes.Buffer)

	if _, err := io.Copy(tmp, f1); err != nil {
		return nil
	}
	s := strings.ToUpper(tmp.String())

	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}

	return nil
}

// CapitalizerExample example of using Capitalizer
func CapitalizerExample() error {
	f1, err := os.Create("file1.txt")
	if err != nil {
		return err
	}
	if _, err := f1.Write([]byte(`this is some long
	enoying string
	`)); err != nil {
		return err
	}

	f2, err := os.Create("file2.txt")
	if err != nil {
		return err
	}

	if err := Capitalizer(f1, f2); err != nil {
		return err
	}

	if err := f1.Close(); err != nil {
		return err
	}
	if err := f2.Close(); err != nil {
		return err
	}

	if err := os.Remove("file1.txt"); err != nil {
		return err
	}

	if err := os.Remove("file2.txt"); err != nil {
		return err
	}

	return nil
}
