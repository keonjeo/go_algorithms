//Package sorts a package for demonstrating sorting algorithms in Go
package sorts

func shellSort(data []int) []int {
	for d := int(len(data) / 2); d > 0; d /= 2 {
		for i := d; i < len(data); i++ {
			for j := i; j >= d && data[j-d] > data[j]; j -= d {
				data[j], data[j-d] = data[j-d], data[j]
			}
		}
	}
	return data
}
