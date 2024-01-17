package lz77


func (lz77 *LZ77) Decode(dictionary []CTuple) []byte {
	//maxSearch := lz77.Config.SearchBuffer
	resultArray := make([]byte,0)
	var defaultByte byte
	for _, tup := range dictionary {

		toAdd := resultArray[len(resultArray) - tup.Offset:len(resultArray) - tup.Offset + tup.Length]
		for _, b := range toAdd {
			resultArray = append(resultArray, b)
		}

		
		if tup.NextByte != defaultByte{
			resultArray = append(resultArray, tup.NextByte)

		}
			

	}
	return resultArray
}
