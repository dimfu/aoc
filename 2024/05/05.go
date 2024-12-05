package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func mustatoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pages := make(map[int][]string)
	ptproduce := [][]int{}

	eopages := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) == 0 {
			eopages = true
			continue
		}
		if !eopages {
			pnots := strings.Split(txt, "|")
			x, y := mustatoi(pnots[0]), mustatoi(pnots[1])
			pages[x] = append(pages[x], txt)
			pages[y] = append(pages[y], txt)
		} else {
			seq := strings.Split(txt, ",")
			seqInt := make([]int, len(seq))
			for i, num := range seq {
				seqInt[i] = mustatoi(num)
			}
			ptproduce = append(ptproduce, seqInt)
		}
	}

	invalidOrderingRules := [][]int{}
	validOrderingRules := [][]int{}

	for _, pt := range ptproduce {
		invalid := false
		for pass := 0; pass < len(pt); pass++ {
			for i := 0; i < len(pt)-1; i++ {
				curr, next := pt[i], pt[i+1]
				if pg, exists := pages[curr]; exists {
					pair_reverse := fmt.Sprintf("%v|%v", next, curr)
					for _, pgstr := range pg {
						if pgstr == pair_reverse {
							pt[i], pt[i+1] = pt[i+1], pt[i]
							invalid = true
						}
					}
				}
			}
		}
		if !invalid {
			validOrderingRules = append(validOrderingRules, pt)
		} else {
			invalidOrderingRules = append(invalidOrderingRules, pt)
		}
	}

	ans1, ans2 := 0, 0
	for _, pages := range validOrderingRules {
		mid := len(pages) / 2
		ans1 += pages[mid]
	}
	for _, pages := range invalidOrderingRules {
		mid := len(pages) / 2
		ans2 += pages[mid]
	}

	fmt.Println(ans1, ans2)
}
