package mapTasks

/*
Дана мапа map[int]int, где ключи и значения — целые числа. Напиши функцию, которая
находит все пары ключ-значение, такие, что их сумма уникальна.
Верни результат в виде мапы, где ключ — это сумма, а значение — массив пар [ключ, значение].
*/

func UniqueSums(values map[int]int) map[int][]int {
	uniqueCheck := make(map[int]bool)

	for key, value := range values {
		if uniqueCheck[key+value] {
			uniqueCheck[key+value] = false
		} else {
			uniqueCheck[key+value] = true
		}
	}

	uniqueResult := make(map[int][]int)

	for key, value := range values {
		if uniqueCheck[key+value] {
			uniqueResult[key+value] = []int{key, value}
		}
	}

	return uniqueResult
}
