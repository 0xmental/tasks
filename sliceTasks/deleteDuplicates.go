package tasks

/*
Напишите функцию, которая удаляет дубликаты из массива
и возвращает новый массив с уникальными значениями.
*/

func unique(arr []int) []int {
	unqMap := make(map[int]bool)
	for _, value := range arr {
		if unqMap[value] { // Эту проверку можно опустить, но я решил оставить первоначальный вариант решения.
			continue
		}
		unqMap[value] = true
	}

	unqArr := make([]int, 0, len(unqMap))
	for i, _ := range unqMap {
		unqArr = append(unqArr, i)
	}
	return unqArr
}

/*
После того как решил эту задачу, разбирал ее через гпт
Он написал прикольное решение, но к нему я сам не дошел
  for _, value := range arr {
      if !unqMap[value] {
          unqMap[value] = true
          unqArr = append(unqArr, value)
      }
  }
*/
