package algo

//BubbleSort sorts a given array using bubble sort
func BubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j+1] < array[j] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
	return array
}

//SelectionSort sorts a given array using selection sort
func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		if i == len(array)-1 {
			break
		}
		smallestItem := i
		for j := i + 1; j < len(array); j++ {
			if array[smallestItem] > array[j] {
				smallestItem = j
			}
		}
		array[i], array[smallestItem] = array[smallestItem], array[i]
	}
	return array
}
