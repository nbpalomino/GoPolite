** Golang First Project **

This is just a very basic started project for Golang.
It's a very _polite_ library for salute and introduce people.

´´´go
package main

import (
    "fmt"
    "github.com/nbpalomino/polite"
)

func main() {
    fmt.Println(polite.Salute("Nick"))
    fmt.Println(polite.Introduce("Nick Palomino"))
}

´´´
