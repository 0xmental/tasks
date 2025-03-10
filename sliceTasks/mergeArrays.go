package tasks

/*
Напишите функцию, которая принимает два массива и возвращает новый массив,
состоящий из элементов обоих массивов.
*/

func merge(firstArr []int, secondArr []int) []int {

	result := make([]int, len(firstArr)+len(secondArr))
	copy(result, firstArr)
	copy(result[len(firstArr):], secondArr)

	return result
}
