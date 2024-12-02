package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	x, y := []int{}, []int{}
	occurs := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), "   ")
		xint, err := strconv.Atoi(c[0])
		if err != nil {
			log.Fatal(err)
		}
		yint, err := strconv.Atoi(c[1])
		if err != nil {
			log.Fatal(err)
		}
		occurs[yint]++
		x = append(x, xint)
		y = append(y, yint)
	}

	slices.Sort(x)
	slices.Sort(y)

	t_distance := 0.0
	s_score := 0

	if len(x) == len(y) {
		for i := range x {
			if oc, exists := occurs[x[i]]; exists {
				s_score += x[i] * oc
			}
			t_distance += math.Abs(float64(x[i] - y[i]))
		}
	}

	fmt.Println(s_score)
}
