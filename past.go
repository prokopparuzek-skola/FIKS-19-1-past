// https://www.geeksforgeeks.org/count-divisors-n-on13/
package main

import "fmt"
import "os"
import "io"
import "math"

const FILE = "prvocisla/primes.txt"
const ODM = int(1e4)
const LIMIT = 1e12

func isPrime(n int, isPrimeArray map[int]bool) bool {
	if n == 1 {
		return false
	}
	if isPrimeArray[n] == true {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var file *os.File
	var T int
	var primes []int
	var isPrimeArray map[int]bool
	var err error
	var buf int

	fmt.Scanf("%d", &T)
	file, _ = os.Open(FILE)
	primes = make([]int, 0)
	isPrimeArray = make(map[int]bool)
	for err != io.EOF {
		_, err = fmt.Fscanf(file, "%d", &buf)
		primes = append(primes, buf)
		isPrimeArray[buf] = true
	}

	for i := 1; i <= T; i++ {
		var N int
		var ans int = 1
		fmt.Scan(&N)
		if N > LIMIT {
			for j := 0; j <= T-i; j++ {
				fmt.Println("O velky Tung")
			}
			return
		}
		for j := 0; int(math.Pow(float64(primes[j]), float64(3))) <= N; j++ {
			cnt := 1
			for N%primes[j] == 0 {
				N /= primes[j]
				cnt++
			}
			ans *= cnt
			if isPrimeArray[N] == true {
				break
			}
		}
		if isPrime(N, isPrimeArray) {
			ans *= 2
		} else if math.Sqrt(float64(N))-float64(int(math.Sqrt(float64(N)))) == 0 {
			if isPrime(int(math.Sqrt(float64(N))), isPrimeArray) {
				ans *= 3
			}
		} else if N != 1 {
			ans *= 4
		}
		ans--
		fmt.Println(ans)
	}
}
