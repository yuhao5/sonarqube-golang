package main

import (
	"fmt"

	"gitlab.kenny.com/demo/math"
)

func main() {
	m := new(math.Math)
	m.A = 1
	m.B = 0
	fmt.Printf("math divide result: %+v\n", m.Divide())
	mathMap, _ := m.ToMap()
	fmt.Printf("mathMap: %+v\n", mathMap)
}
