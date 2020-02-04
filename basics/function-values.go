package main

import (
    "fmt"
    "math"
)

func compute(a, b float64, fn func(float64, float64) float64) float64 {
    return fn(a, b)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }

    fmt.Println(compute(3, 4, hypot))
    fmt.Println(compute(3, 4, math.Pow))
}
