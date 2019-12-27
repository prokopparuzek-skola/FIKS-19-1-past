package main

import "fmt"

const MAX = 1e6

func main() {
	var sito []bool
	sito = make([]bool, MAX)
	sito[0] = true
	for i, t := range sito {
		n := i + 1
		i += n
		if !t {
			for i < len(sito) {
				sito[i] = true
				i += n
			}
		}
	}
	for i, t := range sito {
		if !t {
			fmt.Printf("%d\n", i+1)
		}
	}
}
