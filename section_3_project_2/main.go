package main

import "fmt"

/*
Here are the instructions for the second project in Section 3:

1: Create a new package main
2: In the main func, create a slice of ints that goes from 0-10
3: Iterate through the slice and for each element print whether or not
they are even or odd

*/

func main() {
	nums := []int{}

	for i := 0; i < 11; i++ {
		nums = append(nums, i)
	}

	fmt.Printf("nums: %v\n", nums)
	
	for num := range nums {
		if num % 2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}