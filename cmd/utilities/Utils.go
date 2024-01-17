package utilities

import (
	"os"
)

func GetSrcContent(path string) ([]byte, *error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, &err
	} else {
		return content, nil
	}
}



func SplitEvenly(n int) (int, int) {
	// Calculate the two split values
	half := n / 2
	remainder := n % 2

	// Adjust the split if there's a remainder
	return half + remainder, half
}


func SliceWithPadding(slice []byte, start, end int) []byte {
    // Ensure start index is within bounds
    if start < 0 {
        start = 0
    }

    // Ensure end index is within bounds
    if end > len(slice) {
        end = len(slice)
    }

    // Calculate the length of the resulting slice
    length := end - start
    if length < 0 {
        length = 0
    }

    // Create a new slice with the specified length and fill it with 0's
    result := make([]byte, length)
    copy(result, slice[start:end])

    return result
}
