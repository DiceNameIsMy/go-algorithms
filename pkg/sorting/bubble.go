package sorting

func BubbleSort(data []int) []int {
	var sorted bool
	var swapped bool

	for !sorted {
		swapped = false
		for i := 0; i < len(data)-1; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
	}

	return data
}
