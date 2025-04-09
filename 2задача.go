package main

import (
	"fmt"
	"math"
)

func main() {
	s, p, d := getRactangleData(4, 5)
	fmt.Println("Площадь:", s)
	fmt.Println("Периметр:", p)
	fmt.Println("Диагональ:", d)
}
func getRactangleData(a, b int) (s1, p2 int, d3 float64) {
	s1 = a * b
	p2 = (a + b) * 2
	d3 = math.Sqrt(float64(a*a + b*b))
	return
}
