package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WorkWithTemp() error {
	println(os.TempDir())

	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}
	println(t)

	defer os.RemoveAll(t)

	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}
	fmt.Println(tf.Name())

	return nil
}
