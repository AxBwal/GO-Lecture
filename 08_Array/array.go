package main

import "fmt"

func main() {
	//int -> zero
	//string -> space blank
	//bool -> false   in starting

	// if nothing declraed then this value will show 

	//zeroed values get add in the int type array

	// var nums[4] int

	// nums[0]=1
	// nums[1]=2
	// nums[2]=3
	// nums[3]=4

	// fmt.Println(nums)

	//so in case of boolean it gives me false in all starting
	// var values[3] bool

	// fmt.Println(values)

	// so here it comes with blank spaces if you dont give anything
	// var names[3] string
	// names[1]="golang"
	// fmt.Println(names)

	//to declare array in single line
	nums := [3]int{1, 2, 3}
	fmt.Println(nums)
	fmt.Println("The size is ", len(nums))

	//2d array
	nums2d := [2][2]int{{1, 2}, {2, 1}}
	fmt.Println(nums2d)

	//if the size is fixed then use array by which memory optimization can be good
	//constant time access

	// The size that we can defined through len property

}
