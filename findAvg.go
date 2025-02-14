package tasks

/*
Напишите функцию, которая вычисляет среднее значение
элементов массива и возвращает его.
*/

func avg(arr []int) float64 {
	var sum int

	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}
