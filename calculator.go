package main

import "fmt"

func main() {
    var a, b float64
    var op string

    fmt.Print("Enter calculation (e.g., 10 + 5): ")
    fmt.Scan(&a, &op, &b)

    if op == "+" {
        fmt.Println("Result:", a+b)
    } else if op == "-" {
        fmt.Println("Result:", a-b)
    } else if op == "*" {
        fmt.Println("Result:", a*b)
    } else if op == "/" {
        fmt.Println("Result:", a/b)
    } else {
        fmt.Println("Invalid operator")
    }
}
