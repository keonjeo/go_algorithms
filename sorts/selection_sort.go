//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func selectionSort(data []int) []int {

	for i := 0; i < len(data); i++ {
		min := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[min] {
				min = j
			}
		}

		tmp := data[i]
		data[i] = data[min]
		data[min] = tmp
	}
	return data
}
