//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func bubbleSort(data []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(data)-1; i++ {
			if data[i+1] < data[i] {
				data[i+1], data[i] =  data[i], data[i+1]
				swapped = true
			}
		}
	}
	return data
}
