package main

import "fmt"

// this func does'nt need to be exported to be used in same package
func externalFunc() {
	fmt.Println("This is an external function call")
}
