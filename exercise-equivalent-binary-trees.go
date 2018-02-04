package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
    var walker func(*tree.Tree)
    walker = func(*tree.Tree) {
        if t.Left != nil {
            walker(t.Left)
        }

        ch <- t.Value

        if t.Right != nil {
            walker(t.Right)
        }
    }

    walker(t)
    close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
    c1 := make(chan int)
    c2 := make(chan int)

    go Walk(t1, c1)
    go Walk(t2, c2)

    for {
        n1, ok1 := <-c1
        n2, ok2 := <-c2

        if n1 != n2 || ok1 != ok2 {
            return false
        }

        if !ok1 {
            break
        }
    }

    return true
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(5), tree.New(5)))
    fmt.Println(Same(tree.New(2), tree.New(2)))
}
