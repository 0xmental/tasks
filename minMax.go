package tasks

/*
Создайте функцию, которая находит минимальное и максимальное
 значение в массиве.
*/

func minMax(arr []int) (int, int) {
	min, max := arr[0], arr[0]

	for _, value := range arr {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return min, max
}
