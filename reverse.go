package tasks

/*
Реализуйте функцию, которая принимает массив
и возвращает его в обратном порядке.
*/

func reverse(arr []int) []int {
	j := len(arr) - 1

	for i := 0; i < j; i++ {
		arr[j], arr[i] = arr[i], arr[j]
		j--
	}

	return arr
}
