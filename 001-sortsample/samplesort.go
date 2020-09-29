package main

import (
    "fmt"
    "sort"
)

func main() {
    
    xi := []int{10, 4, 6, 8, 1, 34, 67, 32, 7, 0}
    fmt.Printf("Slice Before Sorting :%v\n", xi)
    searchInts := sort.SearchInts(xi, 10)
    fmt.Println(searchInts)
    // sort int slice
    sort.Ints(xi)
    fmt.Printf("Slice After Sorting :%v\n", xi)
    
    si := []string{"hello", "new", "divya", "adersh", "prayag", "kaira"}
    fmt.Printf("Slice Before Sorting :%v\n", si)
    // sort string slice
    sort.Strings(si)
    fmt.Printf("Slice After Sorting :%v\n", si)
    
    // returns the index
    ints := sort.SearchInts(xi, 10)
    fmt.Println(ints)
    
}
