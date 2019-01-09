package swagger

import (
	"fmt"
	"testing"
)

func TestReadFileSuccess(t *testing.T) {
	result, err := ReadFile("../../asset/apartment-api.yaml")
	if err == nil {
		fmt.Println("result: " + result)
	} else {
		t.Fatalf("failed test %#v", err)
	}
}

func TestReadFileFailed(t *testing.T) {
	result, err := ReadFile("dummy")
	if err == nil {
		fmt.Println(result)
		t.Fatal("failed test")
	}
}
