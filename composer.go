package main

import (
	"math"
)

// generate is a function that produces odd numbers until the limit max.
// It send a channel signal when the limit is reached
func generate(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 2 {
		ch <- i
	}
	ch <- -1 // signal che abbiamo finito
}

// filter is a function that copies the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		if i%prime != 0 {
			out <- i
		}
	}
	out <- -1
}

// calcPrimeFactors return the prime factors of a number. It is derived from a slightly modified version of
// sieve.go in Go source distribution.
func calcPrimeFactors(number_to_factorize int) []int {
	rv := []int{}
	ch := make(chan int)                 // Create a new channel.
	go generate(number_to_factorize, ch) // Start generate() as goroutine.
	for prime := <-ch; (prime != -1) && (number_to_factorize > 1); prime = <-ch {
		for number_to_factorize%prime == 0 {
			number_to_factorize = number_to_factorize / prime
			rv = append(rv, prime)
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	return rv
}

// getBaseAndHeight takes as argument the an array containing the prime factors of a number
// and gives back widht and height of the rectangle
func getBaseAndHeight(prime_factors []int) (bool, int, int) {
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

// calculateRectangle takes as parameter the an integer, that is the total of images that should
// compose the rectangle. If the number is a square number, the method return the side of the square as base
// and height, skipping the calculation. If it's not, the function calls itself recursively, removing one element each time
// is not possible to find out a rectangle
func calculateRectangle(rct map[string]int) map[string]int {
	is_square, side := isASquare(rct["area"])
	if is_square {
		rct["height"], rct["base"] = int(side), int(side)
	} else {
		factors := calcPrimeFactors(rct["area"])
		founded, base, height := getBaseAndHeight(factors)
		rct["base"], rct["height"] = base, height
		if founded == false {
			rct["area"] = rct["area"] - 1
			rct["skipped"] = rct["skipped"] + 1
			calculateRectangle(rct)
		}
	}
	return rct
}

// isASquale check if a number is a square number, return the side
func isASquare(number_to_square int) (bool, float64) {
	side := math.Sqrt(float64(number_to_square))
	side_without_decimals := float64(int(side))
	return ((side - side_without_decimals) == 0), side
}

// calculatePositions takes as parameter the dimension of the desidered thumb that will compose the rectangle
// and the list of images to merge. It calculates the exact postion of each image in the final
// rectangle
func calculatePositions(rect map[string]int, images []string, thumb_width int, thumb_height int) map[string][2]int {
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
