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

//const LIMIT = 1152921504606846976 // 2^60
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

func isPrime(n *big.Int, isPrimeArray *map[uint]bool) bool {
	if n.Cmp(big.NewInt(1)) == 0 {
		return false
	}
	if (*isPrimeArray)[uint(n.Uint64())] {
		return true
	} else if n.Cmp(big.NewInt(134217728)) <= 0 {
		return false
	}
	if big.NewInt(0).Mod(n, big.NewInt(2)).Cmp(&big.Int{}) == 0 || big.NewInt(0).Mod(n, big.NewInt(3)).Cmp(&big.Int{}) == 0 {
		return false
	}
	for i := big.NewInt(5); big.NewInt(0).Mul(i, i).Cmp(n) <= 0; i.Add(i, big.NewInt(6)) {
		if big.NewInt(0).Mod(n, i).Cmp(&big.Int{}) == 0 || big.NewInt(0).Mod(n, big.NewInt(0).Add(i, big.NewInt(2))).Cmp(&big.Int{}) == 0 {
			return false
		}
	}
	return true
}

func isPrimeProbably(n *big.Int, isPrimeArray *map[uint]bool) bool {
	if n.Cmp(big.NewInt(1e12)) <= 0 {
		return isPrime(n, isPrimeArray)
	}
	if big.NewInt(0).Mod(n, big.NewInt(2)).Cmp(&big.Int{}) == 0 || big.NewInt(0).Mod(n, big.NewInt(3)).Cmp(&big.Int{}) == 0 {
		return false
	}

	for i := 0; i < ACC; i++ {
		cmd := exec.Command("dc")
		cmd.Stdin = strings.NewReader(fmt.Sprintf("%d %d^%d%%p\n", random(), big.NewInt(0).Sub(n, big.NewInt(1)), n))
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
		_, err = fmt.Fscan(file, &buf)
		if err == io.EOF {
			break
		}
		primes = append(primes, buf)
		isPrimeArray[buf] = true
	}

	for i := uint(1); i <= T; i++ {
		var N *big.Int
		N = new(big.Int)
		var ans *big.Int = big.NewInt(1)
		fmt.Scan(N)
		//fmt.Fprintln(os.Stderr, N)
		for j := uint(0); big.NewInt(0).Exp(big.NewInt(int64(primes[j])), big.NewInt(3), nil).Cmp(N) <= 0; j++ {
			cnt := uint(1)
			for big.NewInt(0).Cmp(big.NewInt(0).Mod(N, big.NewInt(int64(primes[j])))) == 0 {
				//fmt.Fprintln(os.Stderr, primes[j])
				N.Div(N, big.NewInt(int64(primes[j])))
				cnt++
				//fmt.Fprintf(os.Stderr, "%d:%v\n", N, isPrimeArray[N])
			}
			ans.Mul(ans, big.NewInt(int64(cnt)))
			if isPrimeProbably(N, &isPrimeArray) {
				break
			}
		}
		if isPrimeProbably(N, &isPrimeArray) {
			ans.Mul(ans, big.NewInt(2))
			//fmt.Fprintln(os.Stderr, "Prime")
		} else if big.NewFloat(0).Sqrt(big.NewFloat(0).SetInt(N)).IsInt() {
			var pow *big.Float = new(big.Float)
			var Ipow *big.Int = new(big.Int)
			//fmt.Fprintln(os.Stderr, N)
			pow = pow.SetInt64(N.Int64())
			pow.Sqrt(pow)
			Ipow, _ = pow.Int(nil)
			if isPrimeProbably(Ipow, &isPrimeArray) {
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
