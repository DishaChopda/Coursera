package main

import "fmt"

func main() {

fmt.Printf("Enter float value\n")
var a float64
fmt.Scanln(&a)
var b int=int(a)
fmt.Printf("Float Truncated to Integer is %d\n",b)


}
	