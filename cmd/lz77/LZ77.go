package lz77

import (
	"github.com/rodrigosemicolon/gompress/cmd/utilities"
	"strconv"
	"encoding/gob"
	"encoding/binary"
	"os"
	"fmt"
)

type CTuple struct {
	Offset   int
	Length   int
	NextByte byte
}

type LZ77 struct {
	SearchBufferSize    int
	LookAheadBufferSize int
}

func LZ77FromWindow(windowLength int) *LZ77 {
	searchSize, lookAheadSize := utilities.SplitEvenly(windowLength)
	return &LZ77{
		SearchBufferSize:    searchSize,
		LookAheadBufferSize: lookAheadSize,
	}
}

func NewLZ77FromBuffers(searchSize, lookAheadSize int) *LZ77 {
	return &LZ77{
		SearchBufferSize:    searchSize,
		LookAheadBufferSize: lookAheadSize,
	}
}

func ExtractTuples(content []byte) ([]CTuple, *error){
	return []CTuple{}, nil
}

func (t *CTuple) ToString() string{
	return "offset: " +  strconv.Itoa(t.Offset) +  "\tlength: " +  strconv.Itoa(t.Length) +  "\tbyte: " +  strconv.Itoa(int(t.NextByte))	
}

func WriteDstContent(path string, content []CTuple) *error{
	f, err := os.Create(path)
	if err != nil{

		defer f.Close()
		fmt.Println("writedstcontent: ", content)
		err2 := binary.Write(f, binary.LittleEndian, content)
		
		
		if err2 != nil{
			return &err2
		}
	}
	return nil
}


func WriteCTuplesToFile(filename string, ctuples []CTuple) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(ctuples); err != nil {
		return fmt.Errorf("Error encoding CTuples: %v", err)
	}

	return nil
}

func ReadCTuplesFromFile(filename string) ([]CTuple, error) {
	var ctuples []CTuple

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&ctuples); err != nil {
		return nil, fmt.Errorf("Error decoding CTuples: %v", err)
	}

	return ctuples, nil
}
