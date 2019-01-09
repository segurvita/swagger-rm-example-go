package swagger

import (
	"fmt"
	"io/ioutil"
)

// practice of println
func HelloWorld() {
	fmt.Println("Hello World!")
}

// open file
func ReadFile(filepath string) (string, error) {
	// open file
	b, err := ioutil.ReadFile(filepath) // just pass the file name
	if err != nil {
		return "", err
	}

	// convert content to string
	str := string(b)
	fmt.Println(str)

	return str, nil
}
