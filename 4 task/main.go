package main

import "fmt"

func main() {
	var num int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Таблица умножения до %d:\n", num)
	for i := 1; i <= num; i++ {
		for j := 1; j <= num; j++ {
			fmt.Printf("%3d", i*j)
		}
		fmt.Println()
	}
}
