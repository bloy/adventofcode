package main

import (
	"bufio"
)

func init() {
	AddSolution(2, day2)
}

func day2LetterCount(txt string) (seenTwo, seenThree bool) {
	runes := []rune(txt)
	seen := make(map[rune]bool)
	for i := 0; i < len(runes)-1; i++ {
		if seen[runes[i]] {
			continue
		}
		count := 0
		seen[runes[i]] = true
		for j := i + 1; j < len(runes); j++ {
			if runes[i] == runes[j] {
				count++
			}
		}
		if count == 1 {
			seenTwo = true
		} else if count == 2 {
			seenThree = true
		}
		if seenTwo && seenThree {
			break
		}
	}
	return
}

func day2Similar(txt1, txt2 string) (bool, string) {
	var runes1, runes2 []rune
	runes1 = []rune(txt1)
	runes2 = []rune(txt2)
	if len(runes1) != len(runes2) {
		return false, ""
	}
	diffs := 0
	sames := make([]rune, 0, len(runes1)-1)
	for i := 0; i < len(runes2); i++ {
		if runes1[i] == runes2[i] {
			sames = append(sames, runes2[i])
		} else {
			diffs++
		}
		if diffs >= 2 {
			return false, ""
		}
	}
	if diffs == 1 {
		return true, string(sames)
	}
	return false, ""
}

func day2(pr *PuzzleRun) {
	var boxIDs []string
	s := bufio.NewScanner(pr.InFile)
	for s.Scan() {
		boxIDs = append(boxIDs, s.Text())
	}
	if err := s.Err(); err != nil {
		pr.logger.Fatalln("error reading input", err)
	}
	pr.ReportLoad()

	var twoCount, threeCount int
	for _, id := range boxIDs {
		seenTwo, seenThree := day2LetterCount(id)
		if seenTwo {
			twoCount++
		}
		if seenThree {
			threeCount++
		}
	}
	pr.ReportPart("2:", twoCount, "3:", threeCount, "Checksum:", twoCount*threeCount)
	for i := 0; i < len(boxIDs)-1; i++ {
		for j := i + 1; j < len(boxIDs); j++ {
			similar, sames := day2Similar(boxIDs[i], boxIDs[j])
			if similar {
				pr.ReportPart(string(sames))
				return
			}
		}
	}
	pr.ReportPart("Not Found")
}
