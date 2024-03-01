package main

import (
	"testing"
)

func TestSum(t *testing.T) {
    if sum(1, 2) != 3 {
        t.Error("Sum function is not working correctly")
    }
}
