package main

import (
    "fmt"
    "os"

    "github.com/lubovskiy/app"
)

func main() {
    a := app.New()
    _ = a

    fmt.Println(os.Getenv("GOPATH"))
}

