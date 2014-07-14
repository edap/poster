package main

import (
	//"errors"
	//"fmt"
	//"image"
	//"image/jpeg"
	"math"
	//"os"
)

// Generate numbers until the limit max.
// after the 2, all the prime numbers are odd
// Send a channel signal when the limit is reached
func Generate(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 // signal che abbiamo finito
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

// Prime factorization derived from slightly modified version of
// sieve.go in Go source distribution.
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

// func GetBaseAndHeight(prime_factors []int) (bool, []int) {
// 	if len(prime_factors) == 1 {
// 		return false, prime_factors
// 		// } else if len(prime_factors) == 2{
// 	} else {
// 		return true, prime_factors
// 	}
// }

//aggiungere aspect ratio?
// 16/9 = 1.777777777777777
// 4/3 = 1.3333333333333
func GetBaseAndHeight(prime_factors []int) (bool, int, int) {
	if len(prime_factors) == 1 {
		return false, prime_factors[0], prime_factors[0]
	} else if len(prime_factors) == 2 {
		return true, prime_factors[1], prime_factors[0]
	} else {
		last_pos := (len(prime_factors) - 1)
		base := prime_factors[last_pos]
		height := prime_factors[0]
		step_back := last_pos
		for i, val := range prime_factors {
			if i > 0 && (i != (last_pos)) {
				if (height * val) < base {
					height = height * val
				} else {
					step_back -= 1
					base = base * prime_factors[step_back]
				}
			}
		}
		return true, base, height
	}
}

//oppure due routine, una che parte a destra dell'array di integers e una che parte a sinistra, si scambiano continuamente
// il chan che dice a quanto sta il loro prodotto e quando la routinche che calcola l'altezza arriva ad un momento
// in cui il suo totale/il_totale_della base assomiglia all'aspect ratio si ferma.

func CalculateRectangle(rct map[string]int) map[string]int {
	//check se area e' un int
	//aspect_ratio := 1.3333333333333
	is_square, side := IsASquare(rct["area"])
	if is_square {
		rct["area"], rct["height"], rct["base"], rct["skipped"] = rct["area"], int(side), int(side), rct["skipped"]
	} else {
		factors := CalcPrimeFactors(rct["area"])
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

func IsASquare(number_to_square int) (bool, float64) {
	side := math.Sqrt(float64(number_to_square))
	side_without_decimals := float64(int(side))
	return ((side - side_without_decimals) == 0), side
}

func CalculatePositions(rect map[string]int, images []string, thumb_width int, thumb_height int) map[string][2]int {
	x, y := 0, 0
	res := make(map[string][2]int)
	for index, value := range images {
		res[string(value)] = [2]int{x, y}
		x -= thumb_width
		if (index+1)%rect["base"] == 0 {
			x = 0
			y -= thumb_height
		}
	}
	return res
}
