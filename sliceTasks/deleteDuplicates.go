package tasks

/*
Напишите функцию, которая удаляет дубликаты из массива
и возвращает новый массив с уникальными значениями.
*/

func unique(arr []int) []int {
	unqMap := make(map[int]bool)
	var unqArr []int
	for _, value := range arr {
		if !unqMap[value] {
			unqMap[value] = true
			unqArr = append(unqArr, value)
		}
	}

	return unqArr
}
