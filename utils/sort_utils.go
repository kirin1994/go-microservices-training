package utils

func BubbleSort(elements []int) {
	keepRuning := true
	for keepRuning {
		keepRuning = false
		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRuning = true
			}
		}
	}
}
