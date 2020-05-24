package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    z := 1.0
    var beforeZ float64

    for count := 0; ; count++ {
        beforeZ = z
        z -= (z*z - x) / (2 * z)

        if z - beforeZ < 0.0001 {
            return z
        }
    }
}

func main() {
    fmt.Println(Sqrt(2))
}
