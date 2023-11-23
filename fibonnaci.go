package main

import "fmt"

func fibonnaci(f uint) uint {
	if f == 1 || f == 0 {
		return f
	}
	return fibonnaci(f-1) + fibonnaci(f-2)
}

func main() {
	var f uint
	for {
		fmt.Println("Enter an fibonnaci number")
		_, err := fmt.Scan(&f)
		if err == nil {
			break
		}
		fmt.Println("Inalid input, try again")
	}
	fmt.Printf("Fibonnaci number= %v\n", fibonnaci(f))
}
