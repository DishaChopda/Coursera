package main

import(
"fmt"
"bufio"
"os"
"encoding/json"
)

func main(){
var name string
var addr string

person:=make(map[string]string)

read_input:=bufio.NewScanner(os.Stdin)

fmt.Println("Enter name")
read_input.Scan()
name=read_input.Text()

fmt.Println("Enter Address")
read_input.Scan()
addr=read_input.Text()

person["name"]=name
person["addr"]=addr

jsonperson,_:=json.Marshal(person)

fmt.Println("JSON Object is ")
fmt.Printf(string(jsonperson))
}