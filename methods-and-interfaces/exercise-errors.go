package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    f := float64(e)
    return fmt.Sprintf("cannot Sqrt negative number: ", f)
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return x, ErrNegativeSqrt(x)
    }

    return math.Sqrt(x), nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
