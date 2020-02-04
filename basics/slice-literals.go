package main

import "fmt"

func main() {
    q := []int{1, 3, 5, 7, 9}
    fmt.Println(q)

    r := []bool{true, false, true, true, false}
    fmt.Println(r)

    s := []struct{
        i int
        b bool
    }{
        {1, true},
        {3, false},
        {5, true},
        {7, true},
        {9, false},
    }

    fmt.Println(s)
}
