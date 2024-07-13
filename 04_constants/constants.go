package main

import "fmt"

func main() {
	const name = "akhil"

	// once a value is declared as the const value then it can not be changed in GO
	// name ="anshu " ==> you can not re-intilaize the value

	fmt.Println(name)

	const (
		PORT = 3000
		host = "LocalHost"
	)

	// you can also pair up the values in GO

	fmt.Println(PORT, host)
}
