package main

import "fmt"

var name = "name"

func calculateTax(price float32) (float32, float32) {
    return price*0.09, price*0.02
}

func birthday(age *int) {
    if (*age > 50) {
        panic("Too old for this!")
    }
}

func main() {
    defer fmt.Println("This is the end.")
    age := 32
    birthday(&age)
    fmt.Println(age)
    fmt.Printf("The pointer is %v and the value is %v.\n", &age, age)
}
