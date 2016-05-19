package main

import (
    "fmt"
)


func main() {
    fmt.Println("---Array")
    var arr []int
    fmt.Printf("%v %v %T %v\n", len(arr), cap(arr), arr, arr)

    fmt.Println("---Slice")
    // sl := []int{4, 2}
    sl := arr[:]
    fmt.Printf("%T %v\n", sl, sl)

    foo([]int{1, 2, 3}...)

    s := append([]int{1, 2}, []int{4, 5, 6, 7}...)
    fmt.Printf("%T %v\n", s, s)

}

func foo(nums ...int) {
    for i := 0; i < len(nums); i++ {
        fmt.Println(nums[i])
    }
}