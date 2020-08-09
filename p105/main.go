package main

func main() {

	multiply := func(value int, multiplier int) int {
		return value * multiplier
	}

	add := func(value int, additive int) int {
		return value + additive
	}

	// (x * 2 + 1) * 2
	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		println(multiply(add(multiply(v, 2), 1), 2))
	}
}