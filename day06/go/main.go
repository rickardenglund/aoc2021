package main

import (
	"fmt"
	"os"
)

var (
	test  = []int{3, 4, 3, 1, 2}
	input = []int{3, 4, 1, 2, 1, 2, 5, 1, 2, 1, 5, 4, 3, 2, 5, 1, 5, 1, 2, 2, 2, 3, 4, 5, 2, 5, 1, 3, 3, 1, 3, 4, 1, 5, 3, 2, 2, 1, 3, 2, 5, 1, 1, 4, 1, 4, 5, 1, 3, 1, 1, 5, 3, 1, 1, 4, 2, 2, 5, 1, 5, 5, 1, 5, 4, 1, 5, 3, 5, 1, 1, 4, 1, 2, 2, 1, 1, 1, 4, 2, 1, 3, 1, 1, 4, 5, 1, 1, 1, 1, 1, 5, 1, 1, 4, 1, 1, 1, 1, 2, 1, 4, 2, 1, 2, 4, 1, 3, 1, 2, 3, 2, 4, 1, 1, 5, 1, 1, 1, 2, 5, 5, 1, 1, 4, 1, 2, 2, 3, 5, 1, 4, 5, 4, 1, 3, 1, 4, 1, 4, 3, 2, 4, 3, 2, 4, 5, 1, 4, 5, 2, 1, 1, 1, 1, 1, 3, 1, 5, 1, 3, 1, 1, 2, 1, 4, 1, 3, 1, 5, 2, 4, 2, 1, 1, 1, 2, 1, 1, 4, 1, 1, 1, 1, 1, 5, 4, 1, 3, 3, 5, 3, 2, 5, 5, 2, 1, 5, 2, 4, 4, 1, 5, 2, 3, 1, 5, 3, 4, 1, 5, 1, 5, 3, 1, 1, 1, 4, 4, 5, 1, 1, 1, 3, 1, 4, 5, 1, 2, 3, 1, 3, 2, 3, 1, 3, 5, 4, 3, 1, 3, 4, 3, 1, 2, 1, 1, 3, 1, 1, 3, 1, 1, 4, 1, 2, 1, 2, 5, 1, 1, 3, 5, 3, 3, 3, 1, 1, 1, 1, 1, 5, 3, 3, 1, 1, 3, 4, 1, 1, 4, 1, 1, 2, 4, 4, 1, 1, 3, 1, 3, 2, 2, 1, 2, 5, 3, 3, 1, 1}
)

func main() {
	if os.Getenv("PART") == "part1" {
		fmt.Printf("\npart1: %s\n", run(input, 80))
	}

	if os.Getenv("PART") == "part2" {
		fmt.Printf("\npart2: %s\n", run(input, 256))
	}
}

func run(startFishes []int, nRounds int) string {
	f := map[int]int{}
	for _, timer := range startFishes {
		f[timer]++
	}

	for r := 0; r < nRounds; r++ {
		newFish := f[0]

		for i := 0; i < 9; i++ {
			f[i] = f[i+1]
		}

		f[8] = newFish
		f[6] += newFish
	}

	sum := 0
	for _, n := range f {
		sum += n
	}

	return fmt.Sprintf("%d", sum)
}
