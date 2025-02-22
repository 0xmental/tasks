package mapTasks

/*
Даны две мапы map[string]int. Напиши функцию, которая объединяет их,
оставляя для каждого ключа только максимальное значение, если ключ встречается в обеих мапах.
*/

func Merge(first map[string]int, second map[string]int) map[string]int {
	highest, lowest := first, second
	if len(second) >= len(first) {
		highest, lowest = second, first
	}

	result := make(map[string]int)
	for key, value := range highest {
		result[key] = value
	}

	for key, value := range lowest {
		if v, exist := result[key]; !exist || value > v {
			result[key] = value
		}
	}

	return result
}
