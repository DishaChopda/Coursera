package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)


type Persons struct {
    Persons []Person `json:"persons"`
}


type Person struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

func main() {
    // Open our jsonFile
    jsonFile, err := os.Open("person.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened person.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    // we initialize our Users array
    var persons Persons

    
    json.Unmarshal(byteValue, &persons)

    // we iterate through every user within our users array and
    // print out the user Type, their name, and their facebook url
    // as just an example
	
    for i := 0; i < len(persons.Persons); i++ {
        fmt.Println("User Type: " + persons.Persons[i].Type)
        fmt.Println("User Age: " + strconv.Itoa(persons.Persons[i].Age))
        fmt.Println("User Name: " + persons.Persons[i].Name)
        fmt.Println("Facebook Url: " + persons.Persons[i].Social.Facebook)
    }

}