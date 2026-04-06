// utils/math.go
package utils

func Add(a, b int) int {
    return a + b
}

// main.go
package main

import (
    "fmt"
    "myapp/utils"
)

func main() {
    fmt.Println(utils.Add(5, 3))
}
