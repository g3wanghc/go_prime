package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		n, _ := strconv.ParseInt(os.Args[1], 10, 64)
		fmt.Println(largestPrime(int(n)))
	}
}

func largestPrime(n int) int {

	var index_arr = make([]bool, n+1)
	memsetRepeat(index_arr, true)

	var factor = 2
	var exhaust = false

	for !exhaust && float64(factor) <= math.Sqrt(float64(n)) {
		var multiple = 2
		for factor*multiple <= n {
			index_arr[factor*multiple] = false
			multiple++
		}
		factor++
		var next_match = indexTrue(index_arr[factor:])
		if next_match != -1 {
			factor += next_match
		} else {
			exhaust = true
		}
	}

	var l_prime = indexTrueReverse(index_arr)

	if l_prime != -1 {
		return l_prime
	}

	return -1
}

func indexTrueReverse(a []bool) int {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] {
			return i
		}
	}
	return -1
}

func indexTrue(a []bool) int {
	for i, e := range a {
		if e == true {
			return i
		}
	}
	return -1
}

func memsetRepeat(a []bool, v bool) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}
