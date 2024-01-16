package lz77

import "fmt"

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
			maxOffset = j
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
	c.LookAheadBuffer = content[:c.Config.LookAheadBuffer]
	for i < len(content) {
		fmt.Println("search buffer: ", c.SearchBuffer)
		match := FindLongestMatch(c.SearchBuffer, c.LookAheadBuffer)
		compressedData = append(compressedData, match)
		moveFwd := match.Length + 1
		for _, b := range content[i : i+moveFwd] {
			c.SearchBuffer = append(c.SearchBuffer, b)
		}
		if len(c.SearchBuffer)+moveFwd > c.Config.SearchBuffer {
			c.SearchBuffer = c.SearchBuffer[moveFwd:]
		}
		c.LookAheadBuffer = content[i+moveFwd : i+moveFwd+c.Config.LookAheadBuffer]
		i = i + moveFwd

	}
	return compressedData
}
