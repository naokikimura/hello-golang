package main

import (
    "crypto/sha1"
    "fmt"
    "io"
    "math/rand"
)

func main() {
    fmt.Println("Hello, World!")

    fmt.Println(rand.Intn(100))

    password := "L8pgrXvde@CGm8b!mE4"
    fmt.Println(password)

    h := sha1.New()
    io.WriteString(h, "His money is twice tainted:")
    fmt.Printf("% x", h.Sum(nil))
}
