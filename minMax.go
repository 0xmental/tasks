package tasks

import "slices"

/*
Создайте функцию, которая находит минимальное и максимальное
 значение в массиве.
*/

func minMax(arr []int) (int, int) {
	min := slices.Min(arr)
	max := slices.Max(arr)
	return min, max
}

// Не уверен можно ли в задаче использовать пакет slices?
