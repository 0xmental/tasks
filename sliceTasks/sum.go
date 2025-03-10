package tasks

/*
Напишите функцию, которая принимает массив целых чисел
и возвращает сумму всех его элементов.
*/
func sum(arr []int) int {
	var sum int
	
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}
