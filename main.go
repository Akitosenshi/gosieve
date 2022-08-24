package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

var myDict = map[int]int{
	10:  4,
	1e2: 25,
	1e3: 168,
	1e4: 1229,
	1e5: 9592,
	1e6: 78498,
	1e7: 664579,
	1e8: 5761455,
}

func main() {
	const max = 1000000
	var sqr = int(math.Sqrt(max))
	var primes = myDict[max]
	var sieve [max]bool
	for i := 1; i < max; i++ {
		sieve[i] = true
	}
	// setup done
	start := time.Now()
	for i := 1; i < sqr; i++ {
		if sieve[i] {
			for j := (i+1)*(i+1) - 1; j < max; j += i+1 {
				sieve[j] = false
			}
		}
	}
	t := time.Since(start)
	// sieve done
	var results string
	v := make([]string, max, max)
	counted := 0
	startFormat := time.Now()
	for i := 0; i < max; i++ {
		if sieve[i] {
			v = append(v, strconv.Itoa(i+1) + " is prime\n")
			counted++
		} else {
			v = append(v, strconv.Itoa(i+1) + " is not prime\n")
		}
	}
	results = strings.Join(v, "")
	tFormat := time.Since(startFormat)
	fmt.Printf("%scounted: %d\nactual:  %d\n\ncalculated in %s\nformatting this output took %s\n", results, counted, primes, t, tFormat)
}
