package main

import "fmt"

// basic fizz buzz
func fizzBuzz001(start int, end int, fizz int, buzz int) {
	for i := start; i < end; i++ {
		output := ""
		if i%fizz == 0 {
			output = output + "fizz"
		}
		if i%buzz == 0 {
			output = output + "buzz"
		}
		if output == "" {
			output = fmt.Sprintf("%d", i)
		}

		fmt.Println(output)
	}
}

// based on Sieve of Eratosthenes
func fizzBuzz002(start int, end int, fizz int, buzz int) {
	list := make([]string, end-start+1)

	// populate sieve
	for i := 0; i <= end-start; i++ {
		list[i] = fmt.Sprintf("%d", i+start)
	}

	fizzStart := 0
	if fizz != start {
		fizzStart = fizz - start%fizz
	}
	buzzStart := 0
	if buzz == start {
		buzzStart = buzz - start%buzz
	}
	for i := fizzStart; i <= end-start; i += fizz {
		list[i] = "fizz"
	}
	for i := buzzStart; i <= end-start; i += buzz {
		if list[i] == "fizz" {
			list[i] = "fizzbuzz"
		} else {
			list[i] = "buzz"
		}
	}

	for _, v := range list {
		fmt.Printf("%s\n", v)
	}
}

func main() {
	// fizzBuzz001(0, 100, 3, 5)
	fizzBuzz002(6, 100, 3, 5)
}
