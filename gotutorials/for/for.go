//for is Go's only looping construct

package main

import "fmt"

func main() {
	// basic single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + i
	}
	//classic initial/condition/after
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//for infinite loop with break
	for {
		fmt.Println("loop")
		break
	}
	// for loop with conditions and conditional continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
