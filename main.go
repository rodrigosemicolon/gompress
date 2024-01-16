package main

import (
	"flag"
	"fmt"

	"github.com/rodrigosemicolon/gompress/cmd/lz77"
	"github.com/rodrigosemicolon/gompress/cmd/utilities"
)

func quicktest() {
	algFlag := flag.String("alg", "lz77", "insert compression algorithm (lz77 or lz78)")
	fileFlag := flag.String("f", "", "insert path of file to be compressed")
	flag.Parse()
	fmt.Println("hello world", *algFlag, *fileFlag)
	cnt, err := utilities.GetFileContent(*fileFlag)
	if err == nil {
		fmt.Println("content:", string(cnt))
	} else {
		fmt.Println("error:", *err)
	}
	lz77config := lz77.LZ77ConfigFromWindow(10)
	fmt.Println("lz77:", lz77.NewLZ77(*lz77config))
}

func main() {
	quicktest()
}
