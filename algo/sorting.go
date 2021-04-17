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

//InsertionSort sort a given array using insertion sort
func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		if array[0] > array[i] {
			temp := array[i]
			copy(array[1:i+1], array[0:i])
			array[0] = temp
		} else {
			for j := 1; j < i; j++ {
				if array[j-1] <= array[i] && array[j] > array[i] {
					array[j], array[i] = array[i], array[j]
				}
			}
		}
	}
	return array
}
