package main

import "fmt"

func main() {
	genericPrint(1)
	genericPrint("a")
	genericPrint(true)

	printInt(1)
	printFloat(2.0)
	printString("srijan")

	printInterface(10)
	printInterface("srijan")
	printInterface(5.6)

	fmt.Println(multiply(10))
	fmt.Println(multiply(10.1))

	fmt.Println(average([]int{1, 2}))
	fmt.Println(average([]float64{2.0, 2.1, 2.2, 1.9, 1.8, 1.9}))

	fmt.Println(reverse([]int{1, 2, 3, 4, 5}))
	fmt.Println(reverse([]string{"a", "b", "c", "d"}))
}

func printInt(i int) {
	fmt.Println(i)
}

func printFloat(f float64) {
	fmt.Println(f)
}

func printString(s string) {
	fmt.Println(s)
}

func genericPrint[T any](t T) {
	fmt.Println(t)
}

func printInterface(a interface{}) {
	fmt.Println(a)
}

func multiply[T int | float64](a T) T {
	return a * 11
}

type Number interface {
	int
	float64
}

func average[T int | float64](a []T) T {
	var sum T

	length := T(len(a))

	for _, v := range a {
		sum += v
	}

	return sum / length
}

func reverseIntegers(list []int) []int {
	first := 0
	last := len(list) - 1

	for first < last {
		list[first], list[last] = list[last], list[first]
		first++
		last--
	}

	return list
}

func reverse[T int | string](list []T) []T {
	first := 0
	last := len(list) - 1

	for first < last {
		list[first], list[last] = list[last], list[first]
		first++
		last--
	}

	return list
}
