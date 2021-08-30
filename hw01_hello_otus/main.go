package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	var x = "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(x))
}
