package aoc

import (
	"fmt"
	"khhk-adventofcode2024/reader"
	"khhk-adventofcode2024/util"
	"slices"
	"strconv"
	"strings"
)

func Aoc() {
	path := "aoc/5"
	// lines := reader.Lines(path + "/example.txt")
	lines := reader.Lines(path + "/input.txt")

	rules, updates := parse(lines)

	// fmt.Println(rules)
	// fmt.Println(updates)

	correct, incorrect := categorizeUpdates(rules, updates)

	sum := part1(correct)
	fmt.Println("Sum 1: ", sum)

	sum2 := part2(rules, incorrect)
	fmt.Println("Sum 2: ", sum2)
}

/* PART 1 SPECIFICS */

func part1(correctUpdates [][]int) int {
	sum := 0

	for _, update := range correctUpdates {
		sum += getMiddle(update)
	}

	return sum
}

func categorizeUpdates(rules []OrderRule, updates [][]int) (correct [][]int, incorrect [][]int) {
	leftMap := make(map[int][]int)

	for _, rule := range rules {
		leftMap[rule.left] = append(leftMap[rule.left], rule.right)
	}

	for _, update := range updates {
		ok := checkUpdate(leftMap, update)
		if ok {
			correct = append(correct, update)
		} else {
			incorrect = append(incorrect, update)
		}
	}

	return
}

func checkUpdate(leftMap map[int][]int, update []int) bool {
	seen := make(map[int]bool)
	for _, val := range update {
		l, ok := leftMap[val]
		if ok {
			for _, target := range l {
				if seen[target] {
					return false
				}
			}
		}
		seen[val] = true
	}

	return true
}

func getMiddle(update []int) int {
	return update[len(update)/2]
}

/* PART 2 SPECIFICS */

func part2(rules []OrderRule, incorrectUpdates [][]int) int {
	sum := 0

	for _, update := range incorrectUpdates {
		fixed := fixUpdate(rules, update)
		sum += getMiddle(fixed)
	}

	return sum
}

func fixUpdate(rules []OrderRule, update []int) (fixed []int) {
	left, right := makeRuleDicts(rules, update)

	remaining := make([]int, len(update))
	copy(remaining, update)

	for len(remaining) > 0 {
		for i := 0; i < len(remaining); i++ {
			val := remaining[i]
			if isLeftmost(right, val) {
				fixed = append(fixed, val)
				remaining = slices.Delete(remaining, i, i+1)
				delete(right, val)
				deleteFromRightVals(right, left[val], val)
				break
			}
		}
	}

	return
}

func deleteFromRightVals(right map[int][]int, rvals []int, val int) {
	for _, rval := range rvals {
		right[rval] = slices.DeleteFunc(right[rval], func(n int) bool { return n == val })
	}
}

func isLeftmost(right map[int][]int, value int) bool {
	return len(right[value]) == 0
}

func makeRuleDicts(rules []OrderRule, update []int) (left map[int][]int, right map[int][]int) {
	left = make(map[int][]int)
	right = make(map[int][]int)

	updateSet := make(map[int]bool)
	for _, val := range update {
		updateSet[val] = true
	}

	for _, rule := range rules {
		if updateSet[rule.left] && updateSet[rule.right] {
			left[rule.left] = append(left[rule.left], rule.right)
			right[rule.right] = append(right[rule.right], rule.left)
		}
	}

	return
}

/* PARSING AND MODELLING */

func parse(lines []string) (rules []OrderRule, updates [][]int) {
	i := 0
	for len(lines[i]) > 0 {
		rule := parseRule(lines[i])
		rules = append(rules, rule)

		i++
	}

	i++

	for i < len(lines) {
		update := parseUpdate(lines[i])
		updates = append(updates, update)
		i++
	}

	return
}

func parseRule(line string) OrderRule {
	spl := strings.Split(line, "|")
	left, _ := strconv.Atoi(spl[0])
	right, _ := strconv.Atoi(spl[1])
	return OrderRule{left, right}
}

func parseUpdate(line string) []int {
	spl := strings.Split(line, ",")
	return util.Map(spl, func(s string) int {
		n, _ := strconv.Atoi(s)
		return n
	})
}

type OrderRule struct {
	left  int
	right int
}
