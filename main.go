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
	fmt.Println(nilM.GetB())

	var mathList []*math.Math
	for i := 0; i < 10; i++ {
		mathList = append(mathList, new(math.Math))
	}

	for i, m := range mathList {
		m.A = i
		m.B = i*2 + 1
	}

	for _, m := range mathList {
		fmt.Println(m.Divide())
	}

	if 1 == 1 {
		if 2 == 2 {
			fmt.Println("ok")
		}
	}
}
