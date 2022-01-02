package main

import "fmt"

type divtwo struct {
	divbytwo int 
	even bool
}

func half(input int) (divtwo) {
 
	even := false
	var modulus int 

    modulus = input % 2	
	if modulus == 0 {		
		even = true		
	}
    
	var ret divtwo
	ret.divbytwo = input / 2
	ret.even = even 		

	return ret

}

func main () {

	var number int

	fmt.Print("Please enter number: ")
	fmt.Scan(&number)

	fmt.Println(half(number))
}

/*
func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	h, even := half(5)
	fmt.Println(h, even)
}
*/