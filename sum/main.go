package main

import (
	"fmt"

	"math/rand"

	"strconv"

	"github.com/fatih/color"
	//1. go get github.com/fatih/color
	//2. import and use
)

func sum(a int, b int) int {
	return a + b
}

func randInt() int {
	return rand.Intn(100)

}

func main() {
	fmt.Println(color.HiRedString(strconv.Itoa(sum(randInt(), randInt()))))
}
