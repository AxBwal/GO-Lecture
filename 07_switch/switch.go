// package main

// import "fmt"

// func main() {

// 	i := 3

// 	switch i{
// 	case 1:
// 		fmt.Println("one")
// 	case 2:
// 		fmt.Println("Two")
// 	case 3:
// 		fmt.Println("Three")
// 	case 4:
// 		fmt.Println("Four")
// 	default:
// 		fmt.Println("other")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("its weekend")
// 	default:
// 		fmt.Println("its working day")
// 	}

// }


package main

import "fmt"

func whoAmI  (i interface{}){
	switch i.(type){
	case string:
		fmt.Println("string")

	default :
	fmt.Println("unknown type")
	}
	
}


func main (){

	whoAmI(3)
}
