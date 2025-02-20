package tasks

/*
Реализуйте функцию, которая считает количество четных и нечетных
чисел в массиве и возвращает их количество.
*/

func evenOdd(arr []int) (int, int) {
	even, odd := 0, 0

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even, odd
}
