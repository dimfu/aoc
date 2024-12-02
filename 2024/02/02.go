package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(lvls []int) bool {
	increasing, decreasing := true, true
	for i := 0; i < len(lvls)-1; i++ {
		gap := lvls[i+1] - lvls[i]

		if gap > 3 || gap < -3 {
			return false
		}

		if gap > 0 {
			decreasing = false
		} else if gap < 0 {
			increasing = false
		} else {
			decreasing = false
			increasing = false
		}

		if !increasing && !decreasing {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	safe, dampened := 0, 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lvlsStr := strings.Split(scanner.Text(), " ")
		levels := make([]int, len(lvlsStr))

		for i, str := range lvlsStr {
			lvl, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			levels[i] = lvl
		}

		if isSafe(levels) {
			safe++
		}

		for i := 0; i < len(levels); i++ {
			clone := make([]int, len(levels))
			copy(clone, levels)
			if isSafe(append(clone[:i], clone[i+1:]...)) {
				dampened++
				break
			}
		}
	}

	fmt.Println(safe, dampened)
}
