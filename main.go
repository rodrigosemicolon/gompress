package main

import (
	"flag"
	"fmt"

	"github.com/rodrigosemicolon/gompress/cmd/lz77"
	"github.com/rodrigosemicolon/gompress/cmd/utilities"
)

func quicktest() {
	//TODO: THERE SHOULD BE AN INTERFACE SURROUNDING THE COMPRESSORS
	algFlag := flag.String("c", "lz77", "insert compression algorithm (lz77 or lz78)")
	srcFileFlag := flag.String("src", "", "insert path of file to be compressed")
	dstFileFlag := flag.String("dst", "", "insert path of the compressed file")
	operationFlag := flag.String("op", "", "insert operation (encode or decode)")
	
	flag.Parse()
	if *algFlag == "lz77" {
		compressor := lz77.NewLZ77FromBuffers(6, 4)
		fmt.Println("lz77:", compressor)
		if *operationFlag == "encode"{
			fmt.Println("hello world", *algFlag, *srcFileFlag)
			cnt, err := utilities.GetSrcContent(*srcFileFlag)
			if err == nil {
				fmt.Println("content:", string(cnt))
				fmt.Println(cnt)
				resultEncoding := compressor.Encode(cnt)
				fmt.Println("resultencoding: ", resultEncoding)
				wErr := lz77.WriteCTuplesToFile(*dstFileFlag, resultEncoding)
				if wErr != nil {
					fmt.Println(wErr)
				}
			} else {
				fmt.Println("error:", *err)
			}
	
		} else if *operationFlag == "decode" {
			//cnt, err := utilities.GetSrcContent(*srcFileFlag)
			//if err == nil {
				//cntTuples, exErr := lz77.ExtractTuples(cnt)
			cntTuples,exErr := lz77.ReadCTuplesFromFile(*srcFileFlag)

			if exErr == nil {
				resultDecoding := compressor.Decode(cntTuples)
				//utilities.WriteDstContent(*dstFileFlag, resultEncoding)
				fmt.Println(string(resultDecoding))

			} else {
				fmt.Println("error:", exErr)
			}
		}
	}


	//var test_search = []byte{23, 24, 11, 29}
	//var test_look = []byte{23, 24, 11, 29}
	//fmt.Println("search:", test_search)
	//fmt.Println("lookahead:", test_look)
	//fmt.Println(lz77.FindLongestMatch(test_search, test_look))
	
	//fmt.Println(resultEncoding)
	//resultDecoding := compressor.Decode(resultEncoding)
	//fmt.Println(cnt, string(cnt))
	//fmt.Println(resultDecoding, string(resultDecoding))

}

func main() {
	quicktest()
}
