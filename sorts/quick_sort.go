//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

import (
	"math/rand"
)

func quickSort(data []int) []int {

	if len(data) <= 1 {
		return data
	}

	median := data[rand.Intn(len(data))]

	lowPart := make([]int, 0, len(data))
	highPart := make([]int, 0, len(data))
	middlePart := make([]int, 0, len(data))

	for _, item := range data {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}
