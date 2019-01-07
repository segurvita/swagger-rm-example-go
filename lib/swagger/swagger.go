package swagger

import (
	"errors"
	"fmt"
)

// practice of println
func HelloWorld() {
	fmt.Println("Hello World!")
}

// practice of test
func example(code string) (int, error) {
	if code == "hoge" {
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}
