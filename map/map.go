package main

import "fmt"

func main() {
    fmt.Println("Hello")
    var m map[string]int
    m = make(map[string]int)
    m["apple"] = 5
    i := m["apple"]
    j := m["root"]
    n := len(m)
    delete(m, "apple")

     ok := m["apple"]
    fmt.Println(i, n, j, ok)



}
