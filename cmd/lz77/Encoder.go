package lz77

import(
	//"github.com/rodrigosemicolon/gompress/cmd/utilities"
)

func FindLongestMatch(searchBuffer, lookAheadBuffer []byte) CTuple {
	maxLength, maxOffset := 0, 0
	matchLength := 0
	var nextByte byte

	for j := 0; j < len(searchBuffer); j++ {
		matchLength = 0
		for k := 0; (k < len(lookAheadBuffer)) && (k+j < len(searchBuffer)); k++ {
			if searchBuffer[j+k] == lookAheadBuffer[k] {
				matchLength++
				//nextByte = lookAheadBuffer[j+k]
			} else {
				break
			}
		}
		if matchLength > 0 && matchLength >= maxLength {
			maxLength = matchLength
			//maxOffset = j
			maxOffset = len(searchBuffer) - j

		}

	}

	if maxLength == 0 {
		nextByte = lookAheadBuffer[0]
	} else if maxLength < len(lookAheadBuffer) {
		nextByte = lookAheadBuffer[maxLength]
	}

	return CTuple{Offset: maxOffset,
		Length:   maxLength,
		NextByte: nextByte}

}

func (c *LZ77) Encode(content []byte) []CTuple {
	compressedData := make([]CTuple, 0)
	i := 0
	LookAheadBuffer := content[:c.LookAheadBufferSize]
	SearchBuffer := make([]byte, c.SearchBufferSize)
	for i < len(content) {
		//fmt.Println("\n\nsearch buffer: ", SearchBuffer, "\tlookahead buffer: ", LookAheadBuffer)
		match := FindLongestMatch(SearchBuffer, LookAheadBuffer)
		compressedData = append(compressedData, match)
		moveFwd := match.Length + 1
		
		SearchBuffer = append(SearchBuffer, content[i : i+moveFwd]...)
		for len(SearchBuffer) > c.SearchBufferSize {
			SearchBuffer = SearchBuffer[1:]
		}
		i = i + moveFwd
		/*
		if i > len(content){
			break
		} else if i + c.LookAheadBufferSize > len(content){
			diff := (i + c.LookAheadBufferSize) - len(content) -1
			LookAheadBuffer = content[i : i+diff]
		} else{
			*/
		
		print("lookaheadbuffer: ", string(LookAheadBuffer))
		print("content: ", string(content[i:]))
		LookAheadBuffer = content[i : i+c.LookAheadBufferSize]
		//LookAheadBuffer = utilities.SliceWithPadding(LookAheadBuffer, i, i + c.LookAheadBufferSize)	
		//}
		//fmt.Println("match:\n", match.ToString())
	}
	return compressedData
}
