package mapTasks

/*
Дана мапа map[string]int, в которой могут быть разные ключи с одинаковыми значениями.
Напиши функцию, которая возвращает массив всех ключей с максимальным значением.
*/

func MaxKeys(input map[string]int) []string {
	var maxKeys []string
	var maxValue int

	for key, value := range input {
		if value > maxValue {
			maxValue = value
			maxKeys = []string{key}
		}

		if value == maxValue {
			maxKeys = append(maxKeys, key)
		}
	}

	return maxKeys
}
