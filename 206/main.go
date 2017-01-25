package main

import (
	"fmt"
	"math"
)

func IsSquare(n int) bool {
	v := int(math.Sqrt(float64(n)))
	return v*v == n
}

func IsRightForm(n int) bool {
	str := fmt.Sprintf("%d", n)

	if len(str) != 19 {
		return false
	}

	for i := 0; 2*i < 19; i++ {
		r := fmt.Sprintf("%d", (i+1)%10)[0]
		if str[2*i] != r {
			return false
		}
	}

	return true
}

func main() {
	low := 1020304050607080900
	high := 1929394959697989990

	bottom := int(math.Sqrt(float64(low)))
	top := int(math.Sqrt(float64(high)))

	for i := bottom; i <= top; i++ {
		if IsRightForm(i * i) {
			fmt.Println(i, "->", i*i)
			break
		}
	}
}
