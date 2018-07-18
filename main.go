package main

import (
	"fmt"

	"gitlab.kenny.com/demo/math"
)

var Unused int

func main() {
	m := math.New(10, 2)

	result := m.Divide()
	fmt.Printf("%d/%d=%d\n", m.GetA(), m.GetB(), result)

	mathMap, _ := m.ToMap()
	fmt.Printf("mathMap: %+v\n", mathMap)
}