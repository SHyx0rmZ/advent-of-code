package input

import (
	"io/ioutil"
	"os"
)

func ReadInput(path string) ([]byte, error) {
	var f *os.File
	var err error

	if path == "-" {
		f = os.Stdin
	} else {
		f, err = os.OpenFile(path, os.O_RDONLY, 0755)
		if err != nil {
			return nil, err
		}
		defer f.Close()
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return c, nil
}
