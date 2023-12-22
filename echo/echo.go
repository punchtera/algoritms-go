package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	filename := filepath.Base(os.Args[0])
	fmt.Printf("commandName %s.exe \n", filename)

	now := time.Now()
	r := buildString(os.Args[1:])
	fmt.Printf("%s elapsed \n", time.Since(now))

	now1 := time.Now()
	r1 := buildStringPro(os.Args[1:])
	fmt.Printf("%s elapsed \n", time.Since(now1))

	fmt.Println(r)
	fmt.Println(r1)
}

func buildString(sArr []string) string {

	s, sep := "", ""
	for _, arg := range sArr {
		s += sep + arg
		sep = " "
	}

	return s
}

func buildStringPro(sArr []string) string {
	return strings.Join(sArr, " ")
}
