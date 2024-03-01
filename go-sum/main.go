package main

import (
	"fmt"
	"math/rand"
)


func sum(a int, b int) int {
    return a + b
}

func randInt() int {
	return rand.Intn(100)
	
	
}

func main() {
	fmt.Println(sum(randInt(),randInt()))
}