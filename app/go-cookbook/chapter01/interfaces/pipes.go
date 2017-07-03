package interfaces

import (
	"io"
	"os"
)

// PipeExample helps give some more examples of using io
func PipeExample() error {
	r, w := io.Pipe()

	go func() {
		w.Write([]byte("testn"))
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return nil
}
