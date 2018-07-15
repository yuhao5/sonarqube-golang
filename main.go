package main

import (
	"fmt"

	"gitlab.kenny.com/demo/math"
)

var Test int

func main() {
	m := new(math.Math)
	m.A = 1
	m.B = 0
	fmt.Printf("math divide result: %+v\n", m.Divide())
	mathMap, _ := m.ToMap()
	fmt.Printf("mathMap: %+v\n", mathMap)
	var nilM *math.Math
	fmt.Println(nilM.GetA())
	fmt.Println(nilM.GetB)
}
