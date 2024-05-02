package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("Введите число: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Println(formatComputers(input))
}

func formatComputers(n int) string {
	if n%10 == 1 && n%100 != 11 {
		return fmt.Sprintf("%d компьютер", n)
	} else if n%10 >= 2 && n%10 <= 4 && (n%100 < 10 || n%100 >= 20) {
		return fmt.Sprintf("%d компьютера", n)
	}
	return fmt.Sprintf("%d компьютеров", n)
}
