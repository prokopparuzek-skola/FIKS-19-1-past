// https://www.geeksforgeeks.org/count-divisors-n-on13/
package main

import "fmt"
import "os"
import "io"
import "math"
import "math/rand"
import "time"

const FILE = "prvocisla/primes.txt"
const LIMIT = 1152921504606846976 // 2^60
const ACC = 10

func random(n uint) (r uint) {
	for {
		r = uint(rand.Uint64()) % (n - 2)
		if r > 2 {
			break
		}
	}
	return
}

func isPrime(n uint, isPrimeArray map[uint]bool) bool {
	if n == 1 {
		return false
	}
	if isPrimeArray[n] == true {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := uint(134217725); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func isPrimeProbably(n uint, isPrimeArray map[uint]bool) bool {
	if n == 1 {
		return false
	}
	if isPrimeArray[n] == true {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 0; i < ACC; i++ {
		if uint(math.Pow(float64(random(n)), float64(n-1)))%n != 1 {
			return false
		}
	}
	return true
}

func main() {
	var file *os.File
	var T uint
	var primes []uint
	var isPrimeArray map[uint]bool
	var err error
	var buf uint
	rand.Seed(time.Now().Unix())

	fmt.Scanf("%d", &T)
	file, _ = os.Open(FILE)
	primes = make([]uint, 0)
	isPrimeArray = make(map[uint]bool)
	for err != io.EOF {
		_, err = fmt.Fscanf(file, "%d", &buf)
		primes = append(primes, buf)
		isPrimeArray[buf] = true
	}

	for i := uint(1); i <= T; i++ {
		var N uint
		var ans uint = 1
		fmt.Scan(&N)
		//fmt.Fprintln(os.Stderr, N)
		if N > LIMIT || N == 0 {
			for j := uint(0); j <= T-i; j++ {
				fmt.Println("O velky Tung")
			}
			return
		}
		for j := uint(0); uint(math.Pow(float64(primes[j]), float64(3))) <= N; j++ {
			cnt := uint(1)
			for N%primes[j] == 0 {
				N /= primes[j]
				cnt++
			}
			ans *= cnt
			if isPrimeArray[N] == true {
				break
			}
		}
		if isPrimeProbably(N, isPrimeArray) {
			ans *= 2
		} else if math.Sqrt(float64(N))-float64(int(math.Sqrt(float64(N)))) == 0 {
			if isPrimeProbably(uint(math.Sqrt(float64(N))), isPrimeArray) {
				ans *= 3
			}
		} else if N != 1 {
			ans *= 4
		}
		ans--
		fmt.Println(ans)
	}
}
