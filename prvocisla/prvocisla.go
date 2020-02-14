package main

import "fmt"

func main() {
	MAX := 106528682
	var sito []bool
	var i uint64
	var t bool
	sito = make([]bool, MAX)
	sito[0] = true
	for i = 0; i < uint64(len(sito)); i++ {
		t = sito[i]
		n := i + 1
		j := i + n
		if !t {
			for j < uint64(len(sito)) {
				sito[j] = true
				j += n
			}
		}
	}
	for i = 0; i < uint64(len(sito)); i++ {
		t = sito[i]
		if !t {
			fmt.Printf("%d\n", i+1)
		}
	}
}
