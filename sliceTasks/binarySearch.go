package tasks

/*
Реализуйте функцию, которая принимает массив и значение, а затем
возвращает индекс первого вхождения этого значения в массиве.
Если элемента нет, возвращает -1.  (Алгоритм бинарного поиска)
*/

func binarySearch(arr []int, x int) int {
	right := len(arr)
	left := -1

	for right-left > 1 {
		mid := (left + right) / 2

		if arr[mid] >= x {
			right = mid
		} else {
			left = mid
		}
	}

	if arr[right] == x {
		return right
	}

	return -1
}
