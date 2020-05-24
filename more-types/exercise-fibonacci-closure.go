package main

import "fmt"

func fibonacci() func() int {
    fib := []int{0, 1}
    return func() {
        fibLen := len(fib)
        n := fib[fibLen - 1] + fib[fibLen - 2]
        fib = append(fib, n)
        return fib[fibLen-2]
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
