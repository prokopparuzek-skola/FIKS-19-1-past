// https://www.geeksforgeeks.org/count-divisors-n-on13/
package main

import "fmt"
import "os"
import "io"
import "math"
import "math/rand"
import "time"
import "os/exec"
import "strings"

const FILE = "prvocisla/primes.txt"
const LIMIT = 1152921504606846976 // 2^60
const ACC = 5

func random() (r uint) {
	for {
		r = uint(rand.Uint64()) % (32)
		if r > 2 {
			break
		}
	}
	return
}

func isPrime(n uint, isPrimeArray *map[uint]bool) bool {
	if n == 1 {
		return false
	}
	if (*isPrimeArray)[n] {
		return true
	} else if n <= 134217728 {
		return false
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := uint(5); i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func isPrimeProbably(n uint, isPrimeArray *map[uint]bool) bool {
	if n <= uint(1e12) {
		return isPrime(n, isPrimeArray)
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 0; i < ACC; i++ {
		cmd := exec.Command("dc")
		cmd.Stdin = strings.NewReader(fmt.Sprintf("%d %d^%d%%p\n", random(), n-1, n))
		outb, _ := cmd.Output()
		out := string(outb)
		var result uint
		fmt.Sscanf(out, "%d", &result)
		if result != 1 {
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
	for {
		_, err = fmt.Fscanf(file, "%d", &buf)
		if err == io.EOF {
			break
		}
		primes = append(primes, buf)
		isPrimeArray[buf] = true
	}

	for i := uint(1); i <= T; i++ {
		var N uint
		var ans uint = 1
		fmt.Scanf("%d", &N)
		//fmt.Fprintln(os.Stderr, N)
		if N > LIMIT || N == 0 {
			for j := uint(0); j <= T-i; j++ {
				fmt.Println("O velky Tung")
			}
			return
		}
		for j := uint(0); primes[j]*primes[j]*primes[j] <= N; j++ {
			cnt := uint(1)
			for N%primes[j] == 0 {
				//fmt.Fprintln(os.Stderr, primes[j])
				N /= primes[j]
				cnt++
				//fmt.Fprintf(os.Stderr, "%d:%v\n", N, isPrimeArray[N])
			}
			ans *= cnt
			if isPrimeArray[N] {
				break
			}
		}
		if isPrimeProbably(N, &isPrimeArray) {
			ans *= 2
			//fmt.Fprintln(os.Stderr, "Prime")
		} else if math.Sqrt(float64(N))-float64(int(math.Sqrt(float64(N)))) == 0 {
			if isPrime(uint(math.Sqrt(float64(N))), &isPrimeArray) {
				ans *= 3
				//fmt.Fprintln(os.Stderr, "PrimePow")
			}
		} else if N != 1 {
			ans *= 4
			//fmt.Fprintln(os.Stderr, "Multi")
		}
		ans--
		fmt.Println(ans)
	}
}
