package aoc

import (
	"container/list"
	"fmt"
	"khhk-adventofcode2024/reader"
	"strconv"
	"strings"
)

func Aoc() {
	path := "aoc/11"
	// lines := reader.Lines(path + "/example.txt")
	lines := reader.Lines(path + "/input.txt")

	stones := parse(lines)
	//matrix.PrintMatrix(arr)

	sum := part1(linkedlist(stones))
	fmt.Println("Sum 1: ", sum)

	sum2 := part2(stones)
	fmt.Println("Sum 2: ", sum2)
}

/* PART 1 SPECIFICS */

func part1(l *list.List) int {
	for i := 0; i < 25; i++ {
		for e := l.Front(); e != nil; e = e.Next() {
			val := e.Value.(int)
			if val == 0 {
				e.Value = 1
				continue
			}

			s := strconv.Itoa(val)
			if len(s)%2 == 0 {
				left, _ := strconv.Atoi(s[:len(s)/2])
				right, _ := strconv.Atoi(s[len(s)/2:])

				l.InsertBefore(left, e)
				nexte := l.InsertBefore(right, e)
				l.Remove(e)
				e = nexte
				continue
			}

			e.Value = val * 2024
		}
	}

	return l.Len()
}

/* PART 2 SPECIFICS */

func part2(stones []int) int {
	cache := make(map[int]int)
	for _, s := range stones {
		cache[s] = cache[s] + 1
	}

	fmt.Println(cache)

	for i := 0; i < 75; i++ {
		nextcache := make(map[int]int)
		for k, v := range cache {
			if k == 0 {
				nextcache[1] = nextcache[1] + v
				continue
			}

			s := strconv.Itoa(k)
			if len(s)%2 == 0 {
				left, _ := strconv.Atoi(s[:len(s)/2])
				right, _ := strconv.Atoi(s[len(s)/2:])

				nextcache[left] = nextcache[left] + v
				nextcache[right] = nextcache[right] + v
				continue
			}

			nk := k * 2024
			nextcache[nk] = nextcache[nk] + v
		}

		cache = nextcache
	}

	sum := 0
	for _, v := range cache {
		sum += v
	}

	return sum
}

/* MODELLING */

/* PARSING */

func parse(lines []string) (stones []int) {
	for _, s := range strings.Split(lines[0], " ") {
		i, _ := strconv.Atoi(s)
		stones = append(stones, i)
	}

	return
}

func linkedlist(stones []int) *list.List {
	ll := list.New()

	for _, i := range stones {
		ll.PushBack(i)
	}

	return ll
}
