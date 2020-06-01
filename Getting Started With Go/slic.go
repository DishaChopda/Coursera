package main

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
)

func main(){

    var user_int string
    sli := make([]int, 3)

    x := false
    i := 0

    fmt.Printf("The initialized slice contains: ")
    fmt.Print(sli)

    for x == false{

        fmt.Printf("\nPlease enter one integer: ")
        fmt.Scan(&user_int)

        if (strings.HasPrefix(user_int, "X") || strings.HasPrefix(user_int, "x")) {

              // Exit condition
              x = true

        } else if valid_int, err := strconv.Atoi(user_int); err == nil {

              // Condition where valid input is added to slice

              sli = append(sli, valid_int)

              fmt.Printf("Current slice contains: ")
              sort.Ints(sli)
              fmt.Print(sli)

        } else{

            i-- // decrements i for invalid input where nothing is added to slice.
            fmt.Printf("Invalid input. Input must be an integer or 'X' to exit")
        }

        i++

    }
}