package main

import "fmt"

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is par")
    } else {
        fmt.Println("7 is inpar")
    }

    if 8%4 == 0 {
        fmt.Println("8 es divisible entre 4")
    }

    if num := 9; num < 0 {
        fmt.Println(num, "es negativo")
    } else if num < 10 {
        fmt.Println(num, "tiene un digito")
    } else {
        fmt.Println(num, "tiene varios dígitos")
    }
}