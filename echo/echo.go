package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filename := filepath.Base(os.Args[0])
	fmt.Printf("commandName %s.exe \n", filename)

	r := buildString(os.Args[1:])

	fmt.Println(r)
}

func buildString(sArr []string) string {

	s, sep := "", ""
	for i, arg := range sArr {
		s += sep + arg
		sep = " "
		fmt.Printf("[%d] : [%s]\n", i, arg)
	}

	return s
}
