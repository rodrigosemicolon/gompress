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
	lz77config := lz77.NewLZ77ConfigFromBuffers(6, 4)
	compressor := lz77.NewLZ77(*lz77config)
	fmt.Println("lz77:", compressor)
	//var test_search = []byte{23, 24, 11, 29}
	//var test_look = []byte{23, 24, 11, 29}
	//fmt.Println("search:", test_search)
	//fmt.Println("lookahead:", test_look)
	//fmt.Println(lz77.FindLongestMatch(test_search, test_look))
	fmt.Println(cnt)
	resultEncoding := compressor.Encode(cnt)
	fmt.Println(resultEncoding)
	resultDecoding := compressor.Decode(resultEncoding)
	fmt.Println(cnt, string(cnt))
	fmt.Println(resultDecoding, string(resultDecoding))

}

func main() {
	quicktest()
}
