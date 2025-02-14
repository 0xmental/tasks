package tasks

import "slices"

/*
Реализуйте функцию, которая принимает массив
и возвращает его в обратном порядке.
*/

func reverse(arr []int) []int {
	slices.Reverse(arr)
	return arr
}

// Не уверен можно ли в задаче использовать пакет slices?
