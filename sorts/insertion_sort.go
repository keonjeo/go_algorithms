//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func insertionSort(data []int) []int {
	for out := 1; out < len(data); out++ {
		temp := data[out]
		in := out

		for ; in > 0 && data[in-1] >= temp; in-- {
			data[in] = data[in-1]
		}
		data[in] = temp
	}
	return data
}
