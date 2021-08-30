package main

import "golang.org/x/example/stringutil"
import "fmt"

func main() {
	var x = "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(x))
}
