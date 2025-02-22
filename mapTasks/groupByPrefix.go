package mapTasks

/*
Напиши функцию, которая принимает массив строк и группирует их по префиксу заданной длины.
Результат — мапа map[string][]string, где ключи — это уникальные префиксы,
а значения — массивы слов с этим префиксом.
*/

func SortingPrefix(words []string, prefix int) map[string][]string {
	sortedByPrefix := make(map[string][]string)

	for _, v := range words {
		if len(v) >= prefix {
			sortedByPrefix[v[:prefix]] = append(sortedByPrefix[v[:prefix]], v)
		}
	}

	return sortedByPrefix
}
