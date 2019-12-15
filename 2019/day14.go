package main

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func init() {
	AddSolution(14, solveDay14)
}

// Ingredient is a matter compiler ingredient
type Ingredient struct {
	Qty  int
	Name string
}

// Reaction is a matter compiler reaction
type Reaction struct {
	In  []Ingredient
	Out Ingredient
}

func day14oreToProduce(desired Ingredient, reactions map[string]Reaction, inventory map[string]int) int {
	if desired.Name == "ORE" {
		return desired.Qty // scoop it up out of space
	}
	if inventory[desired.Name] >= desired.Qty {
		inventory[desired.Name] -= desired.Qty
		return 0
	}
	need := desired.Qty
	if inventory[desired.Name] > 0 {
		need -= inventory[desired.Name]
		inventory[desired.Name] = 0
	}
	reaction := reactions[desired.Name]
	batches := int(math.Ceil(float64(need) / float64(reaction.Out.Qty)))

	ore := 0
	for _, input := range reaction.In {
		ingr := Ingredient{Name: input.Name, Qty: input.Qty * batches}
		ore += day14oreToProduce(ingr, reactions, inventory)
	}

	produced := batches * reaction.Out.Qty
	inventory[desired.Name] += produced - need
	return ore
}

func solveDay14(pr *PuzzleRun) {
	reactions := make(map[string]Reaction)
	scanner := bufio.NewScanner(pr.InFile)
	for scanner.Scan() {
		sides := strings.Split(scanner.Text(), " => ")
		parts := strings.Split(sides[0], ", ")
		lhs := make([]Ingredient, len(parts))
		for i, part := range parts {
			strs := strings.Split(part, " ")
			num, err := strconv.Atoi(strs[0])
			pr.CheckError(err)
			lhs[i] = Ingredient{Qty: num, Name: strs[1]}
		}
		strs := strings.Split(sides[1], " ")
		num, err := strconv.Atoi(strs[0])
		pr.CheckError(err)
		r := Reaction{In: lhs, Out: Ingredient{Qty: num, Name: strs[1]}}
		reactions[r.Out.Name] = r
	}
	if err := scanner.Err(); err != nil {
		pr.CheckError(err)
	}
	pr.ReportLoad()

	totals := make(map[string]int)
	oreTotal := day14oreToProduce(Ingredient{1, "FUEL"}, reactions, totals)
	pr.ReportPart(oreTotal)

	var trillion = 1000000000000
	start, end := 0, trillion
	guesses := 0
	lastGuess, fuelGuess := 0, 0
	for {
		lastGuess = fuelGuess
		guesses++
		fuelGuess = (end-start)/2 + start
		totals = make(map[string]int)
		requiredOre := day14oreToProduce(Ingredient{Name: "FUEL", Qty: fuelGuess}, reactions, totals)
		if requiredOre > trillion {
			end = fuelGuess
		} else {
			start = fuelGuess
		}
		if fuelGuess == lastGuess {
			break
		}
	}
	pr.ReportPart(fuelGuess, guesses)
}
