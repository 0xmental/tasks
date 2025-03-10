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
