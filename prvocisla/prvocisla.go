package main

import "fmt"
import "os"
import "io"

const MAX = 1e12
const SEEK = 1e8
const FILE = "prv.txt"

func main() {
	sito := make([]bool, SEEK)
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
	var soubor *os.File
	soubor, _ = os.Create(FILE)
	for i, t := range sito {
		if !t {
			str := fmt.Sprintf("%d\n", i+1)
			soubor.WriteString(str)
		}
	}
	soubor.Close()
	for i := 1; i <= MAX/SEEK; i++ {
		var buf int
		var err error
		prvocisla := make([]int, 0)
		soubor, _ = os.Open(FILE)
		for {
			_, err = fmt.Fscanf(soubor, "%d", &buf)
			if err == io.EOF {
				break
			}
			prvocisla = append(prvocisla, buf)
		}
		soubor.Close()
		fmt.Fprintf(os.Stderr, "GO:%d\n", i)
		//sito := make(map[int]bool, SEEK)
		sito := make([]bool, SEEK)
		for _, p := range prvocisla {
			n := p * (i * SEEK / p)
			if n <= i*SEEK {
				n += p
			}
			for n <= i*SEEK+len(sito) {
				index := n - i*SEEK - 1
				sito[index] = true
			}
		}
		soubor, _ = os.OpenFile(FILE, os.O_APPEND|os.O_WRONLY, 0664)
		for j, t := range sito {
			if !t {
				str := fmt.Sprintf("%d\n", j+1+i*SEEK)
				soubor.WriteString(str)
			}
		}
		soubor.Close()
	}
}
