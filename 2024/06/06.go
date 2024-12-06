package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// i cant answer part 2, credits: https://www.youtube.com/watch?v=v96h9BMwrSY
func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gmap := [][]rune{}
	var sr, sc int
	p1, p2 := 0, 0

	y := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		gmap = append(gmap, row)
		for x, prop := range row {
			if prop == '^' {
				sr, sc = y, x
				break
			}
		}
		y++
	}

	R, C := len(gmap), len(gmap[0])

	for o_r := 0; o_r < R; o_r++ {
		for o_c := 0; o_c < C; o_c++ {
			r, c := sr, sc
			d := 0 // 0=up, 1=right, 2=down, 3=left
			seen := make(map[[3]int]bool)
			seenRC := make(map[[2]int]bool)

			for {
				if seen[[3]int{r, c, d}] {
					p2++
					break
				}

				seen[[3]int{r, c, d}] = true
				seenRC[[2]int{r, c}] = true

				dr, dc := [4]int{-1, 0, 1, 0}[d], [4]int{0, 1, 0, -1}[d]
				rr, cc := r+dr, c+dc

				// Check boundaries
				if rr < 0 || rr >= R || cc < 0 || cc >= C {
					if gmap[o_r][o_c] == '#' {
						p1 = len(seenRC)
					}
					break
				}

				// Check obstacle or target point
				if gmap[rr][cc] == '#' || (rr == o_r && cc == o_c) {
					d = (d + 1) % 4
				} else {
					r, c = rr, cc
				}
			}
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
