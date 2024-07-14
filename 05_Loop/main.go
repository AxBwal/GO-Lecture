package main

import "fmt"

func main() {

	//while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println()



	//for loop

	for j := 1; j <= 6; j++ {
		fmt.Println(j)
	}



	//infintie loop
	//  for {

	//  } //never use for infinite loop
	fmt.Println()


	//Range starts from 0 ends to the prev numbers
	for k := range 7 {
		fmt.Println(k)
	}
}

// NoTe :== In GO we Dont have While loop like other Languages
// Instead we have for that to be known as both for while loop and for loop
