package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const testString string = `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`

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
	strs = strings.Split(testString, "\n")
	sort.Strings(strs)
	days := make([]day, 0)
	currentDay := day{}
	currentDate := ""
	currentGuard := 0
	for _, str := range strs {
		fmt.Println(str)
		matches := matcher.FindStringSubmatch(str)
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
				currentDay.events = append(currentDay.events, event{0, false})
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
	fmt.Println("Date        ID     Minute")
	fmt.Println("                   000000000011111111112222222222333333333344444444445555555555")
	fmt.Println("                   012345678901234567890123456789012345678901234567890123456789")
	for _, thisDay := range days {
		fmt.Printf("%s  #%4d  ", thisDay.date, thisDay.guard)
		min := 0
		c := "."
		for _, e := range thisDay.events {
			for ; min < e.minute; min++ {
				fmt.Printf("%s", c)
			}
			if e.sleep {
				c = "#"
			} else {
				c = "."
			}
		}
		for ; min < 60; min++ {
			fmt.Printf("%s", c)
		}
		fmt.Println("")
	}
	return days
}

func runPart1(input []day) {
	fmt.Println("part 1 ")
}

func runPart2(input []day) {
	fmt.Println("part 2 ")
}

func main() {
	input := getInput()
	runPart1(input)
	runPart2(input)
}
