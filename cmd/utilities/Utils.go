package utilities

import "os"

type Queue struct {
	Content  []byte
	Capacity int
	Size     int
}

func NewQueue(capacity int) *Queue {
	return &Queue{Content: make([]byte, capacity),
		Capacity: capacity,
		Size:     0}
}

func (q *Queue) Enqueue(element byte) {
	if q.Size < q.Capacity {
		q.Content = append(q.Content, element)
	} else {
		q.Content = append(q.Content[1:], element)
	}
	q.Size = q.Size + 1
}

func GetFileContent(path string) ([]byte, *error) {
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
