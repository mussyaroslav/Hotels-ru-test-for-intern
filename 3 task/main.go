package main

import (
	"fmt"
)

// Функция isPrime проверяет, является ли число простым.
func isPrime(num int) bool {
	// Если число меньше или равно 1, оно не является простым.
	if num <= 1 {
		return false
	}
	// Проверяем деление числа на все числа от 2 до корня из num.
	for i := 2; i*i <= num; i++ {
		// Если число делится на i без остатка, оно не является простым.
		if num%i == 0 {
			return false
		}
	}
	// Если число не делится нацело ни на одно число до корня из num, оно простое.
	return true
}

// Функция generateSlice генерирует слайс простых чисел в заданном диапазоне.
func generateSlice(min, max int) []int {
	// Создаем пустой слайс для хранения простых чисел.
	slice := []int{}
	// Проходимся по всем числам в заданном диапазоне.
	for num := min; num <= max; num++ {
		// Если число простое, добавляем его в слайс.
		if isPrime(num) {
			slice = append(slice, num)
		}
	}
	// Возвращаем слайс простых чисел.
	return slice
}

func main() {
	var min, max int

	// Запрашиваем у пользователя ввод минимального и максимального чисел.
	fmt.Print("Минимальное число: ")
	_, err := fmt.Scan(&min)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Print("Максимальное число: ")
	_, err = fmt.Scan(&max)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Проверяем, что минимальное значение диапазона меньше или равно максимальному.
	if min > max {
		fmt.Println("Минимальное число должно быть меньше или равно максимальному.")
		return
	}

	// Выводим список простых чисел в заданном диапазоне.
	fmt.Println(generateSlice(min, max))
}
