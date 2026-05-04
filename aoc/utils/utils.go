package utils

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Pow : binary exponentiation
func Pow(b int, e int) int {
	res := 1
	for e > 0 {
		if e&1 == 1 {
			res *= b
		}
		b = b * b
		e >>= 1
	}
	return res
}

// GCD : gcd of two ints
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM : lcm of two ints
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

// Primes : return all prime numbers < n
func Primes(n int) []int {
	sieve := make([]bool, n)
	for i := 2; i < n; i++ {
		sieve[i] = true
	}
	for i := 2; i < n; i++ {
		if sieve[i] && i*i < n {
			for j := i * i; j < n; j += i {
				sieve[j] = false
			}
		}
	}
	var primes []int
	for i := 2; i < n; i++ {
		if sieve[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// Factorize : return prime factors of n, precompute Primes once in scope
func Factorize(n int, primes []int) []int {
	var factors []int

	for _, p := range primes {
		if n%p == 0 {
			factors = append(factors, p)
		}
	}
	return factors
}

// Assert : simple assertion
func Assert(cond bool, str string) {
	if !cond {
		panic(str)
	}
}
