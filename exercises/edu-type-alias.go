package main

import (
	"fmt"
	"math"
	"strings"
)

type MyStr string //using MyStr as an alias for type string
type MyFloat float64   //using MyFloat as an alias for type float64

func (s MyStr) Uppercase() string {
	return strings.ToUpper(string(s))
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
    }
    return float64(f)
}

func main() {
	fmt.Println(MyStr("test").Uppercase()) // TEST

    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs()) // 1.4142135623730951
}