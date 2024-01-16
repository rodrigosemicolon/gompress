package lz77

import lz77 "github.com/rodrigosemicolon/gompress/cmd/LZ77"

func FindLongestMatch(searchBuffer, lookAheadBuffer []byte) lz77.CTuple {
	maxLength := 0
	matchLength, matchOffset := 0, 0
	var nextByte byte
	maxPossible := min(len(searchBuffer), len(lookAheadBuffer))
	i := 0
	for i < maxPossible {
		matchLength = 0
		for j := 0; (i + j) < maxPossible; j++ {
			if searchBuffer[i+j] == lookAheadBuffer[i+j] {
				matchLength++
				if matchLength >= maxLength {
					maxLength = matchLength
					matchOffset = i
				}
			} else {
				i = i + j
				nextByte = lookAheadBuffer[i+j+1]
				break
			}
		}
	}
	return lz77.CTuple{Offset: matchOffset,
		Length:   matchLength,
		NextByte: nextByte}

}

/*
func (c *LZ77) Encode(content []byte) {
	compressedData := ""
	i := 0
	c.LookAheadBuffer = content[i : i+c.Config.LookAheadBuffer]
	for i < len(content) {
		matchLength := 0
		bestMatch := CTuple{}

		c.LookAheadBuffer = content[i+matchLength : i+matchLength+c.Config.LookAheadBuffer]
		c.SearchBuffer = content[i : i+c.Config.LookAheadBuffer]

	}
}
*/
