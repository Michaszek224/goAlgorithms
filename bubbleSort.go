package main

import "fmt"

func main() {
	myA := []int{9, 8, 2, 4, 7, 1, 2, 4, 6, 1, 2, 4, 0, 12, 4, 5}
	fmt.Printf("Unsorted array: %v\n", myA)
	swapCounter := 0
	for i := 0; i < len(myA); i++ {
		for j := 0; j < len(myA)-1; j++ {
			if myA[i] < myA[j] {
				temp := myA[i]
				myA[i] = myA[j]
				myA[j] = temp
				swapCounter += 1
			}
		}
	}
	fmt.Printf("Sorted array: %v\n", myA)
	fmt.Println("How many swaps were done: ", swapCounter)
}
