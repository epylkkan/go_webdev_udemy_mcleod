package main

import "fmt"

func max(input ...int) (int) {
 
	var largest int

	for i, v := range input {
		if v > largest || i == 0 {
			largest = v
		}
	}

	return largest

}

func main () {

	greatest := max(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println(greatest)
}
