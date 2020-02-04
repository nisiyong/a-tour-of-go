package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {1.234, -2.345},
    "Google": {3.456, 4.567},
}

func main() {
    fmt.Println(m)
}
