package bytestrings

import (
	"bytes"
	"io"
	"io/ioutil"
)

// Buffer demonstrates some ways to initialize bytes Buffers
// These buffers implement an io.Reader interface
func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)

	// 1: From raw bytes
	var b = new(bytes.Buffer)
	b.Write(rawBytes)

	// 2: alternative
	b = bytes.NewBuffer(rawBytes)

	// 3: avoiding the internal byte array
	b = bytes.NewBufferString(rawString)

	return b
}

func toString(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
