// Prime factorization derived from slightly modified version of
// sieve.go in Go source distribution.
package main

import (
	"fmt"
	"math"
)

// Generate numbers until the limit max.
// after the 2, all the prime numbers are odd
// Send a channel signal when the limit is reached
func Generate(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 // signal that the limit is reached
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		if i%prime != 0 {
			out <- i
		}
	}
	out <- -1
}

func CalcPrimeFactors(number_to_factorize int) []int {
	rv := []int{}
	ch := make(chan int)                 // Create a new channel.
	go Generate(number_to_factorize, ch) // Start Generate() as goroutine.
	for prime := <-ch; (prime != -1) && (number_to_factorize > 1); prime = <-ch {
		for number_to_factorize%prime == 0 {
			number_to_factorize = number_to_factorize / prime
			rv = append(rv, prime)
		}
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
	return rv
}

// vanno inseriti in questa funzione:
// la logica per la radice quadrata, in modo da salvare un po di cicli
// la logica per le prorporzioni, vedi per esempio i fattori primi di 48
// che danno come basa 12*4 invece che 6*8
func GetBaseAndHeight(prime_factors []int) (bool, int, int) {
	if len(prime_factors) == 1 {
		return false, prime_factors[0], prime_factors[0]
	} else if len(prime_factors) == 2 {
		return true, prime_factors[1], prime_factors[0]
	} else {
		last_pos := (len(prime_factors) - 1)
		width := prime_factors[last_pos]
		height := prime_factors[0]
		step_back := last_pos
		for i, val := range prime_factors {
			if i > 0 && (i != (last_pos)) {
				if (height * val) < width {
					height = height * val
				} else {
					step_back -= 1
					width = width * prime_factors[step_back]
				}
			}
		}
		return true, width, height
	}
}

func IsASquare(number_to_square int) (bool, float64) {
	side := math.Sqrt(float64(number_to_square))
	side_without_decimals := float64(int(side))
	return ((side - side_without_decimals) == 0), side
}

func CalculateRectangle(rct map[string]int) map[string]int {
	//check se area e' un int
	//aspect_ratio := 1.3333333333333
	is_square, side := IsASquare(rct["area"])
	if is_square {
		rct["area"], rct["height"], rct["base"], rct["skipped"] = rct["area"], int(side), int(side), rct["skipped"]
	} else {
		factors := CalcPrimeFactors(rct["area"])
		fmt.Println(factors)
		founded, base, height := GetBaseAndHeight(factors)
		rct["base"], rct["height"] = base, height
		if founded == false {
			rct["area"] = rct["area"] - 1
			rct["skipped"] = rct["skipped"] + 1
			CalculateRectangle(rct)
		}
	}
	return rct
}

// func main() {
// 	res := map[string]int{
// 		"area":    6,
// 		"height":  0,
// 		"base":    0,
// 		"skipped": 0,
// 	}
// 	fmt.Println(CalculateRectangle(res))
// }
