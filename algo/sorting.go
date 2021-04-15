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
