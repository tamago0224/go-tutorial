package main

import "fmt"

type Vertex struct {
    X, Y int
}

func do(i interface{}) {
    switch v := i.(type) {
    case int :
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
        fmt.Printf("I don't know about type %T\n", v)
    }
}

func main() {
    do(21)
    do("hello")

    v := Vertex{3,4}
    do(v)
}
