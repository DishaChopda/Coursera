package main

import "fmt"

func GenDisplaceFn(acc,v0,s0 float64) func (float64) float64{

	fnn:=func(t float64)float64{
		
		return 0.5*acc*t*t+v0*t+s0
				//return 0.5*a*t*t + v*t+ d

	}
	return fnn
}


func main(){

var acc,v0,s0,t float64

fmt.Println("Enter Acceleration")
fmt.Scan(&acc)

fmt.Println("Enter Initial Velocity")
fmt.Scan(&v0)

fmt.Println("Enter Initial Displacement")
fmt.Scan(&s0)

fn:=GenDisplaceFn(acc,v0,s0)

fmt.Println("Enter Time")
fmt.Scan(&t)

fmt.Println("Displacement:",fn(t))

}