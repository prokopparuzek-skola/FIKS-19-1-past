// https://www.geeksforgeeks.org/count-divisors-n-on13/
package main

import (
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

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
	var err error
	var buf *big.Int
	rand.Seed(time.Now().Unix())

	fmt.Scanf("%d", &T)
	file, _ = os.Open(FILE)
	primes = make([]uint, 0)
	for {
		_, err = fmt.Fscan(file, &buf)
		fmt.Println(buf)
		if err == io.EOF {
			break
		}
		primes = append(primes, uint(buf.Int64()))
	}

	for i := uint(1); i <= T; i++ {
		var N *big.Int
		var ans *big.Int = big.NewInt(1)
		fmt.Scanf("%d", &N)
		//fmt.Fprintln(os.Stderr, N)
		if N.Cmp(big.NewInt(LIMIT)) == 1 || N.Cmp(&big.Int{}) == 0 {
			for j := uint(0); j <= T-i; j++ {
				fmt.Println("O velky Tung")
			}
			return
		}
		for j := uint(0); big.NewInt(0).Exp(big.NewInt(int64(primes[j])), big.NewInt(3), nil).Cmp(N) <= 0; j++ {
			cnt := uint(1)
			for big.NewInt(0).Cmp(big.NewInt(0).Mod(N, big.NewInt(int64(primes[j])))) == 0 {
				//fmt.Fprintln(os.Stderr, primes[j])
				N.Div(N, big.NewInt(int64(primes[j])))
				cnt++
				//fmt.Fprintf(os.Stderr, "%d:%v\n", N, isPrimeArray[N])
			}
			ans.Mul(ans, big.NewInt(int64(cnt)))
			if N.ProbablyPrime(ACC) {
				break
			}
		}
		if N.ProbablyPrime(ACC) {
			ans.Mul(ans, big.NewInt(2))
			//fmt.Fprintln(os.Stderr, "Prime")
		} else if big.NewFloat(0).Sqrt(big.NewFloat(0).SetInt(N)).IsInt() {
			var pow *big.Float
			var Ipow *big.Int
			pow.Sqrt(big.NewFloat(0).SetInt(N)).Int(nil)
			Ipow, _ = pow.Int(nil)
			if Ipow.ProbablyPrime(ACC) {
				ans.Mul(ans, big.NewInt(3))
				//fmt.Fprintln(os.Stderr, "PrimePow")
			}
		} else if N.Cmp(big.NewInt(1)) != 0 {
			ans.Mul(ans, big.NewInt(4))
			//fmt.Fprintln(os.Stderr, "Multi")
		}
		ans.Sub(ans, big.NewInt(1))
		fmt.Println(ans)
	}
}
