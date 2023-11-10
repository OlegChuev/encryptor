package filesystem

import (
	"io/ioutil"
)

func WriteToFile(path string, result []byte) (err error) {
	err = ioutil.WriteFile(path, result, 0644)

	return
}
