package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	fmt.Printf("Command %s \n", os.Args[0])
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Printf("[%d] : [%s]\n", i, arg)
	}
	fmt.Println(s)
}
