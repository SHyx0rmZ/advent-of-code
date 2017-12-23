package input

import (
	"io"
	"io/ioutil"
	"os"
)

func OpenInputFile(path string) (io.ReadCloser, error) {
	if path == "-" {
		return struct {
			io.Reader
			io.Closer
		}{
			os.Stdin,
			ioutil.NopCloser(nil),
		}, nil
	}

	return os.OpenFile(path, os.O_RDONLY, 0644)
}

func ReadInput(path string) ([]byte, error) {
	f, err := OpenInputFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return c, nil
}
