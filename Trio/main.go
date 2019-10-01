package main

import "fmt"
import "math"

func main() {
	var x, sin, i float64
	fmt.Scan(&x)
	var degree float64 = x
	x = 360 - x%360
	x = 22 * x / (7.0 * 180)
	fmt.Println(degree, x)
	var count int = 1
	for i = 1; i <= 1000; i += 2 {
		if count%2 == 1 {
			sin += math.Pow(x, i) / fact(i)
		} else {
			sin -= math.Pow(x, i) / fact(i)
		}
		count++
	}
	cos := math.Pow(1-math.Pow(sin, 2), 0.5)
	var tan, sec, cot, cosec float64
	if degree == 90 {
		tan = sin / 0.0
		sec = cos / 0.0
		cot = 0
		cosec = 1
	}else {
		cosec = 1.0/sin
		tan = sin/cos 
		cot = cos/sin 
		sec = 1/cos
	}
	if degree/90 == 2{
		fmt.Println(degree,degree/90)
		cos = -cos
		tan = -tan
		sec = -sec
		cot = -cot
	}
	fmt.Println(sin, cos, tan, cosec, sec, cot)
}
func fact(i float64) float64 {
	if i == 0 || i == 1 {
		return 1
	} else {
		return i * fact(i-1)
	}
}
