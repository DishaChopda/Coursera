package main
import(
"fmt"
"bufio"
"os"
"strings"
) 

type animal struct{
food string
locomtion string
spoken string
}

func (a *animal)Eat(){
fmt.Println(a.food)
}

func (a *animal)Move(){
fmt.Println(a.locomtion)
}

func (a *animal)Speak(){
fmt.Println(a.spoken)
}

func Store(s []string){
	obj:=new(animal)
	
	switch s[0]{
		case "cow":
			obj.food = "grass"
			obj.locomtion = "walk"
			obj.spoken = "moo"
			
		case "bird":
			obj.food = "worms"
			obj.locomtion = "fly"
			obj.spoken = "peep"
	
		case "snake":
			obj.food = "mice"
			obj.locomtion = "slither"
			obj.spoken = "hsss"
	}
	
	switch s[1]{
		case "eat":
			obj.Eat()
			
		case "move":
			obj.Move()
		
		case "speak":
			obj.Speak()
	
	}
}

func main(){
	input:=bufio.NewReader(os.Stdin)
	for{
		fmt.Println(">")
		a,_:=input.ReadString('\n')
		s:=strings.Split(strings.TrimSpace(a)," ")
		
		Store(s)
	}
}
