package main
import "fmt"

func Swap(s []int, i int){
	s[i],s[i+1] = s[i+1],s[i]
}

func BubbleSort(a []int){
	
	n:=len(a)
	for i:=0; i < n-1; i++{
	
		for j:=0; j < n-i-1; j++{
		
			if a[j] > a[j+1]{
			
				Swap(a,j)
			}
		}
	}
}


func main(){
	fmt.Printf("Enter 10 Integers")
	arr:=make([]int,10)
	for i:=0;i<10;i++{
		fmt.Scan(&arr[i])
	}
	BubbleSort(arr)
	fmt.Println(arr)
}
