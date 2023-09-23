package main

import (
	"fmt"
	"math/rand"
)

func RandomColor() string {
	return fmt.Sprintf("rgb(%v,%v,%v)", RGB(), RGB(), RGB())
}

func RGB() int {
	return rand.Intn(256)
}
