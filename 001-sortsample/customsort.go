package main

import (
    "fmt"
    "sort"
)

func main() {
    vento := car{brand: "vw", model: "vento", cc: 1600}
    polo := car{brand: "vw", model: "polo", cc: 1198}
    jetta := car{brand: "vw", model: "jetta", cc: 2000}
    tiguan := car{brand: "vw", model: "tiguan", cc: 2000}
    taigun := car{brand: "vw", model: "taiguan", cc: 1000}
    troc := car{brand: "vw", model: "troc", cc: 1500}
    
    cars := []car{vento, polo, jetta, taigun, tiguan, troc}
    fmt.Println("Before Sorting :", cars)
    // conversion of struct []cars to struct []byCC and then sorting
    sort.Sort(byCC(cars))
    fmt.Println("After Sorting By CC", cars)
    
    sort.Sort(byModel(cars))
    fmt.Println("After Sorting By Model", cars)
    
}

type car struct {
    brand, model string
    cc           int
}

func (c car) String() string {
    return fmt.Sprintf("%s-%s: %d\n", c.brand, c.model, c.cc)
}

type byCC []car
type byModel []car

// Both these types are implementing Interface.Sort
func (b byModel) Len() int {
    return len(b)
}

// Less basic comparison of which value is Less
func (b byModel) Less(i, j int) bool {
    return b[i].model < b[j].model
}

func (b byModel) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}

func (b byCC) Len() int {
    return len(b)
}

func (b byCC) Less(i, j int) bool {
    return b[i].cc < b[j].cc
}

func (b byCC) Swap(i, j int) {
    b[i], b[j] = b[j], b[i]
}
