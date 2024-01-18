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
	srcFileFlag := flag.String("src", "test2.txt", "insert path of file to be compressed")
	dstFileFlag := flag.String("dst", "test2.lz77", "insert path of the compressed file")
	operationFlag := flag.String("op", "encode", "insert operation (encode or decode)")
	windowSizeFlag := flag.Int("winsize", 10, "insert window size for the compression algorithm (default: 10)")
	searchSizeFlag := flag.Int("ssize", -1, "insert search buffer size for the compression algorithm (default: -1) if this is provided, lookahead buffer size must also be provided")
	lookaheadSizeFlag := flag.Int("lsize", -1, "insert lookahead buffer size for the compression algorithm (default: -1)")

	flag.Parse()

	if *algFlag == "lz77" {
		var compressor lz77.LZ77
		if *searchSizeFlag != -1 && *lookaheadSizeFlag != -1 {
			compressor = *lz77.NewLZ77FromBuffers(*searchSizeFlag, *lookaheadSizeFlag)
		} else {
			compressor = *lz77.NewLZ77FromWindow(*windowSizeFlag)
		}

		if *operationFlag == "encode" {
			fmt.Println("using ", *algFlag, "to encode", *srcFileFlag)
			cnt, err := utilities.GetSrcContent(*srcFileFlag)
			if err == nil {
				resultEncoding := compressor.Encode(cnt)
				wErr := lz77.WriteCTuplesToFile(*dstFileFlag, resultEncoding)
				if wErr != nil {
					fmt.Println(wErr)
				}
			} else {
				fmt.Println("error:", *err)
			}

		} else if *operationFlag == "decode" {
			fmt.Println("using ", *algFlag, "to decode", *srcFileFlag)

			cntTuples, exErr := lz77.ReadCTuplesFromFile(*srcFileFlag)

			if exErr == nil {
				resultDecoding := compressor.Decode(cntTuples)
				fmt.Println(string(resultDecoding))

			} else {
				fmt.Println("error:", exErr)
			}
		}
	}

}

func main() {
	quicktest()
}
