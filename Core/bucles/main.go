package main

func main() {
	// for
	for i := 0; i < 10; i++ {
		println(i)
	}

	// while
	k := 10
	for k > 0 {
		println(k)
		k--
	}

	// for con arreglos
	arr := [3]int{1, 2, 3}
	for index, value := range arr {
		println(index, value)
	}
}
