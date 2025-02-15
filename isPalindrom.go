package tasks

// Создайте функцию, которая проверяет, является ли массив палиндромом

func isPalindrom(arr []int) bool {

	for i := 0; i > len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}

/*
Альтернативное решение, про пакет рефлект узнал через гпт
Решение придумал сам, но решил за основу взять вариант без пакета
	copyArr := append([]int{}, arr...)
	slices.Reverse(copyArr)
	return reflect.DeepEqual(arr, copyArr)
*/
