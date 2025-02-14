package tasks

/*
Напишите функцию, которая принимает два массива и возвращает новый массив,
состоящий из элементов обоих массивов.
*/

func mergeArr(arr1 []int, arr2 []int) []int {
	result := append(arr1, arr2...)
	return result
}
