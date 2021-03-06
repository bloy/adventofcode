package main

import (
	"regexp"
	"strconv"
	"strings"
)

const armyRegex = `^(.*):$`
const groupRegex = `^(\d+) units each with (\d+) hit points (?:\((.*)\) )?with an attack that does (\d+) (\w+) damage at initiative (\d+)$`

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func parseInput(str string) []*Group {
	lines := strings.Split(str, "\n")
	groups := make([]*Group, 0)
	armyRe := regexp.MustCompile(armyRegex)
	groupRe := regexp.MustCompile(groupRegex)
	army := ""
	for _, line := range lines {
		if line == "" {
			continue
		}
		match := armyRe.FindStringSubmatch(line)
		if len(match) > 0 {
			army = match[1]
			continue
		}
		match = groupRe.FindStringSubmatch(line)
		g := NewGroup(army, atoi(match[1]), atoi(match[2]), atoi(match[4]), atoi(match[6]), match[5], match[3])
		groups = append(groups, g)
	}
	return groups
}

const inputStr = `Immune System:
2749 units each with 8712 hit points (immune to radiation, cold; weak to fire) with an attack that does 30 radiation damage at initiative 18
704 units each with 1890 hit points with an attack that does 26 fire damage at initiative 17
1466 units each with 7198 hit points (immune to bludgeoning; weak to slashing, cold) with an attack that does 44 bludgeoning damage at initiative 6
6779 units each with 11207 hit points with an attack that does 13 cold damage at initiative 4
1275 units each with 11747 hit points with an attack that does 66 cold damage at initiative 2
947 units each with 5442 hit points with an attack that does 49 radiation damage at initiative 3
4319 units each with 2144 hit points (weak to bludgeoning, fire) with an attack that does 4 fire damage at initiative 9
6315 units each with 5705 hit points with an attack that does 7 cold damage at initiative 16
8790 units each with 10312 hit points with an attack that does 10 fire damage at initiative 5
3242 units each with 4188 hit points (weak to cold; immune to radiation) with an attack that does 11 bludgeoning damage at initiative 14

Infection:
1230 units each with 11944 hit points (weak to cold) with an attack that does 17 bludgeoning damage at initiative 1
7588 units each with 53223 hit points (immune to bludgeoning) with an attack that does 13 cold damage at initiative 12
1887 units each with 40790 hit points (immune to radiation, slashing, cold) with an attack that does 43 fire damage at initiative 15
285 units each with 8703 hit points (immune to slashing) with an attack that does 60 slashing damage at initiative 7
1505 units each with 29297 hit points with an attack that does 38 fire damage at initiative 8
191 units each with 24260 hit points (immune to bludgeoning; weak to slashing) with an attack that does 173 cold damage at initiative 20
1854 units each with 12648 hit points (weak to fire, cold) with an attack that does 13 bludgeoning damage at initiative 13
1541 units each with 49751 hit points (weak to cold, bludgeoning) with an attack that does 62 slashing damage at initiative 19
3270 units each with 22736 hit points with an attack that does 13 slashing damage at initiative 10
1211 units each with 56258 hit points (immune to slashing, cold) with an attack that does 73 bludgeoning damage at initiative 11`

const testStr = `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4`
