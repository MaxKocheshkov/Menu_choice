package main

import "fmt"

func main() {
	menu := [6]string{"Water", "Burger", "Cake", "Soup", "Soda", "Fries"}

	var pos int
	fmt.Scanln(&pos)

	for i := 0; i < len(menu); i++ {
		if pos == i {
			fmt.Println(menu[i])
			break
		}
		if pos > len(menu) {
			fmt.Println("Invalid choice")
			break
		}

	}

}
