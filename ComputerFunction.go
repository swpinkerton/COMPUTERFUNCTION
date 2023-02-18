package main

import "fmt"

// COMPUTERFUNCTION
//Simply XORs the two inputs
func COMPUTERFUNCTION(a uint, b uint) uint {
	var total uint
	total = a ^ b
	return total
}

func main() {
	inputs := [3][3]uint{
		{10, 20, 30},
		{17, 35, 50},
		{61, 233, 212}}

	var i int
	for i = 0; i < 3; i++ {
		fmt.Printf("COMPUTERFUNCTION = %d : Answer = %d \n", COMPUTERFUNCTION(inputs[i][0], inputs[i][1]), inputs[i][2])
	}
}
