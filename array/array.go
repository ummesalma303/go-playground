package main

import "fmt"

func main() {
    fmt.Println("Hello")

    //  int array 
    var a [4]int
    a[2] = 1
    i := a[2]
    fmt.Println(a, i)
    // string 
    b := [5]string{"Penn", "Teller"}
    fmt.Println(b)
    // slice
    letters := []string{"a", "b", "c", "d"}
    fmt.Println(letters)

}
