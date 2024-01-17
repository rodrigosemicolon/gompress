package lz77

import (
	"github.com/rodrigosemicolon/gompress/cmd/utilities"
	"strconv"
)

type CTuple struct {
	Offset   int
	Length   int
	NextByte byte
}

type LZ77 struct {
	//Dictionary map[string]CTuple
	SearchBuffer    []byte
	LookAheadBuffer []byte
	Config          LZ77Config
}

type LZ77Config struct {
	SearchBuffer    int
	LookAheadBuffer int
}

func LZ77ConfigFromWindow(windowLength int) *LZ77Config {
	searchSize, lookAheadSize := utilities.SplitEvenly(windowLength)
	return &LZ77Config{
		SearchBuffer:    searchSize,
		LookAheadBuffer: lookAheadSize,
	}
}

func NewLZ77ConfigFromBuffers(searchSize, lookAheadSize int) *LZ77Config {
	return &LZ77Config{
		SearchBuffer:    searchSize,
		LookAheadBuffer: lookAheadSize,
	}
}

func NewLZ77(config LZ77Config) *LZ77 {
	return &LZ77{
		SearchBuffer:    make([]byte, 0), //config.SearchBuffer),
		LookAheadBuffer: make([]byte, config.LookAheadBuffer),
		Config:          config,
	}
}

func (t *CTuple) ToString() string{
	return "offset: " +  strconv.Itoa(t.Offset) +  "\tlength: " +  strconv.Itoa(t.Length) +  "\tbyte: " +  strconv.Itoa(int(t.NextByte))	
}