package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	fmt.Println("Hello, world!")

	var x int = 5
	y := 7
	var sum int
	sum = x + y
	fmt.Println(sum)

	var a [5]int
	a[2] = 7
	fmt.Println(a)

	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	slice := []int{5, 4, 3, 2, 1}
	slice = append(slice, 13)
	fmt.Println(slice)

	vertices := make(map[string]int)

	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12

	delete(vertices, "square")
	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for index, value := range b {
		fmt.Println("index:", index, "value:", value)
	}

	for key, value := range vertices {
		fmt.Println("index:", key, "value:", value)
	}
	
	fmt.Println(mySumFunc(2, 2))

	sqrtResult, err := sqrt(25)

	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println(sqrtResult)
	}
}

func mySumFunc (x int, y int) int {
	return x + y
}

func sqrt (x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}