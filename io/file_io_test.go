package io

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	file := ReadFile("E://Account.rec0")
	fmt.Println(string(file))
}
