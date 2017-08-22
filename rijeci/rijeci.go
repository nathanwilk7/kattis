package main

import "fmt"

func rijeci (k int) (int, int) {
	var curAs int
	numAs, numBs := 1, 0
	for i := 0; i < k; i++ {
		curAs = numAs
		numAs = numBs
		numBs += curAs
	}
	return numAs, numBs
}
func main () {
	var k int
	fmt.Scanln(&k)
	numAs, numBs := rijeci(k)
	fmt.Println(numAs, numBs)
}
