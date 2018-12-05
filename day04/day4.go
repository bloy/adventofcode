package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type event struct {
	minute int
	sleep  bool
}

type day struct {
	date   string
	guard  int
	events []event
}

func getInput() []day {
	matcher := regexp.MustCompile(`^\[(\d{4}-\d{2}-\d{2}) \d{2}:(\d{2})\] (.*)$`)
	guardMatcher := regexp.MustCompile(`Guard #(\d+) begins shift`)
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(content), "\n")
	sort.Strings(strs)
	days := make([]day, 0)
	currentDay := day{}
	currentDate := ""
	currentGuard := 0
	for _, str := range strs {
		fmt.Println(str)
		matches := matcher.FindStringSubmatch(str)
		if len(matches) == 0 {
			continue
		}
		date := matches[1]
		minute, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		log := matches[3]
		matches = guardMatcher.FindStringSubmatch(log)
		if len(matches) == 2 {
			num, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			currentGuard = num
		} else {
			if date != currentDate {
				if currentDate != "" {
					days = append(days, currentDay)
				}
				currentDay = day{date, currentGuard, make([]event, 0)}
				currentDate = date
			}
			sleep := false
			if log == "falls asleep" {
				sleep = true
			}
			currentDay.events = append(currentDay.events, event{minute, sleep})
		}
	}
	days = append(days, currentDay)
	return days
}

func runPart1(input []day) {
	guardTotals := make(map[int]int)
	guardMinutes := make(map[int]map[int]int)
	fmt.Println("\n")
	fmt.Println("Date        ID     Minute")
	fmt.Println("                   000000000011111111112222222222333333333344444444445555555555")
	fmt.Println("                   012345678901234567890123456789012345678901234567890123456789")
	for _, d := range input {
		fmt.Printf("%s  #%4d  ", d.date, d.guard)
		guard := d.guard
		min := 0
		sleeping := false
		if guardMinutes[guard] == nil {
			guardMinutes[guard] = make(map[int]int)
		}
		for _, e := range d.events {
			diff := e.minute - min
			if sleeping {
				guardTotals[guard] = guardTotals[guard] + diff
			}
			for i := min; i < e.minute; i++ {
				if sleeping {
					guardMinutes[guard][i] = guardMinutes[guard][i] + 1
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			if e.sleep {
				sleeping = true
			} else {
				sleeping = false
			}
			min = e.minute
		}
		if sleeping {
			diff := 60 - min
			guardTotals[guard] = guardTotals[guard] + diff
		}
		for i := min; i < 60; i++ {
			if sleeping {
				guardMinutes[guard][i] = guardMinutes[guard][i] + 1
				fmt.Print("$")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	maxTotal := 0
	maxGuard := 0
	for guard, total := range guardTotals {
		if total > maxTotal {
			maxGuard = guard
			maxTotal = total
		}
	}
	maxMinute := -1
	maxTotal = 0
	for min, total := range guardMinutes[maxGuard] {
		if total > maxTotal {
			maxMinute = min
			maxTotal = total
		}
	}
	fmt.Println("part 1 ", maxGuard, maxMinute, maxGuard*maxMinute)
}

func runPart2(input []day) {
	// this time it's minute -> guard -> count
	var guardMinutes [60]map[int]int
	for _, d := range input {
		guard := d.guard
		min := 0
		sleeping := false
		for _, e := range d.events {
			for i := min; i < e.minute; i++ {
				if guardMinutes[i] == nil {
					guardMinutes[i] = make(map[int]int)
				}
				if sleeping {
					guardMinutes[i][guard] = guardMinutes[i][guard] + 1
				}
			}
			sleeping = e.sleep
			min = e.minute
		}
		if sleeping {
			for i := min; i < 60; i++ {
				guardMinutes[i][guard] = guardMinutes[i][guard] + 1
			}
		}
	}
	maxTotal := -1
	maxMinute := -1
	maxGuard := -1
	for i := 0; i < 60; i++ {
		for guard, total := range guardMinutes[i] {
			if total > maxTotal {
				maxTotal = total
				maxMinute = i
				maxGuard = guard
			}
		}
	}
	fmt.Println("part 2 ", maxTotal, maxMinute, maxGuard, maxMinute*maxGuard)
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
