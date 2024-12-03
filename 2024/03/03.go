package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	do_mul  bool
	cleared bool
}

func mustatoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func mul(x, y string) int {
	return mustatoi(x) * mustatoi(y)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	enabled := true

	ans1 := 0
	ans2 := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

		for _, mulmatch := range re.FindAllStringSubmatchIndex(text, -1) {
			x := string(text[mulmatch[2]:mulmatch[3]])
			y := string(text[mulmatch[4]:mulmatch[5]])

			ans1 += mul(x, y)

			re_ins := regexp.MustCompile(`do(?:n't)?\(\)`)
			instruction := re_ins.FindAllString(text[0:mulmatch[5]], -1)

			if len(instruction) > 0 {
				ins := instruction[len(instruction)-1]
				if ins == "do()" {
					enabled = true
					ans2 += mul(x, y)
				} else {
					enabled = false
				}
			} else {
				if enabled {
					ans2 += mul(x, y)
				}
			}
		}
	}
	fmt.Println(ans1, ans2)
}
