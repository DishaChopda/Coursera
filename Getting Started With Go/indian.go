package main
 
import( 
"fmt"
"strings"

)
 
func main() {
        fmt.Print("Enter String: ")   
        var first string    
        fmt.Scanln(&first)                  
   
		pre:=false
		suf:=false
		con:=false
		
		pre=strings.HasPrefix(first,"i")
		suf=strings.HasSuffix(first,"n")
		con=strings.Contains(first,"a")

		
		if pre{
		 if suf{
			 if con{
				fmt.Print("Found")
			 }else{
		fmt.Print("Not Found ")}
		 }else{
		fmt.Print("Not Found ")}
		}else{
		fmt.Print("Not Found ")}
}