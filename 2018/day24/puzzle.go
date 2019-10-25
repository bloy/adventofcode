package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type fights map[*group]*group
type targets map[*group]bool

type dataType []*group

func (data dataType) Info(detail bool) string {
	b := strings.Builder{}
	c := make(dataType, len(data))
	copy(c, data)
	dataBy(func(g1, g2 *group) bool {
		if g1.army == g2.army {
			return g1.power() > g2.power()
		}
		return g1.army < g2.army
	}).Sort(c)
	army := ""
	group := 0
	for i := range c {
		if !c[i].inCombat() {
			continue
		}
		if c[i].army != army {
			fmt.Fprintf(&b, "\n%s:\n", c[i].army)
			army = c[i].army
			group = 0
		}
		group++
		c[i].group = group
		fmt.Fprintf(&b, "%v\n", c[i].Info(detail))
	}
	return b.String()
}

func (data dataType) String() string {
	return data.Info(true)
}

func (data dataType) Copy() dataType {
	newData := make(dataType, len(data))
	for i, g := range data {
		newData[i] = g.Copy()
	}
	return newData
}

func (data dataType) boost(amt int) dataType {
	data = data.Copy()
	for i := range data {
		if data[i].army == "Immune System" {
			data[i].damage = data[i].damage + amt
		}
	}
	return data
}

func (data dataType) armyStats() map[string]int {
	stats := make(map[string]int)
	for i := range data {
		if data[i].inCombat() {
			army := data[i].army
			if _, ok := stats[army]; !ok {
				stats[army] = 0
			}
			stats[army] = stats[army] + data[i].units
		}
	}
	return stats
}

func (data dataType) fightStep(debug bool) dataType {
	allGroups := data.Copy()
	if debug {
		fmt.Print(allGroups.Info(false))
	}
	dataBy(func(g1, g2 *group) bool {
		if g1.army == g2.army {
			g1p, g2p := g1.power(), g2.power()
			if g1p == g2p {
				return g1.init > g2.init
			}
			return g1p > g2p
		}
		return g1.army < g2.army
	}).Sort(allGroups)
	// target selection phase
	fighters := make(fights)
	targetSet := make(targets)
	for i := range allGroups {
		target := allGroups[i].bestTarget(allGroups, targetSet, debug)
		if target != nil {
			targetSet[target] = true
			fighters[allGroups[i]] = target
		}
	}

	// attack phase
	attackerList := make(dataType, len(fighters))
	i := 0
	for g := range fighters {
		attackerList[i] = g
		i++
	}
	dataBy(func(g1, g2 *group) bool {
		return g1.init > g2.init
	}).Sort(attackerList)
	for _, attacker := range attackerList {
		defender := fighters[attacker]
		dmg := attacker.damageVs(defender)
		units := defender.applyDamage(dmg)
		if debug {
			fmt.Printf("%s group %d attacks defending group %d, killing %d units\n",
				attacker.army, attacker.group, defender.group, units)
		}
	}

	return allGroups
}

func (data dataType) runFight(debug bool) (army string, units int) {
	groups := data
	stats := groups.armyStats()
	for len(stats) > 1 {
		groups = groups.fightStep(debug)
		stats = groups.armyStats()
	}
	for a := range stats {
		army, units = a, stats[a]
	}
	if debug {
		fmt.Print("\n----------------------------\n\n")
	}
	return
}

type dataBy func(g1, g2 *group) bool

func (f dataBy) Sort(data dataType) {
	sorter := &dataSorter{dataType: data, by: f}
	sort.Sort(sorter)
}

type dataSorter struct {
	dataType
	by dataBy
}

func (data *dataSorter) Len() int { return len(data.dataType) }
func (data *dataSorter) Swap(i, j int) {
	data.dataType[i], data.dataType[j] = data.dataType[j], data.dataType[i]
}
func (data *dataSorter) Less(i, j int) bool { return data.by(data.dataType[i], data.dataType[j]) }

type stringSet map[string]bool

func (s stringSet) words() []string {
	keys := make([]string, len(s))
	i := 0
	for w := range s {
		keys[i] = w
		i++
	}
	return keys
}

type group struct {
	group, units, hp, damage, init int
	immune, weak                   stringSet
	army, damageType               string
}

func (g *group) Copy() *group {
	newg := &group{
		units:      g.units,
		hp:         g.hp,
		damage:     g.damage,
		init:       g.init,
		army:       g.army,
		damageType: g.damageType,
		immune:     make(stringSet),
		weak:       make(stringSet),
		group:      0,
	}
	for k, v := range g.immune {
		newg.immune[k] = v
	}
	for k, v := range g.weak {
		newg.weak[k] = v
	}
	return newg
}

func (g *group) Info(detail bool) string {
	if detail {
		return g.String()
	}
	return fmt.Sprintf("Group %d contains %d units", g.group, g.units)
}

func (g *group) String() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "%d units each with %d hit points", g.units, g.hp)
	if len(g.immune) > 0 || len(g.weak) > 0 {
		fmt.Fprint(&b, " (")
		if len(g.immune) > 0 {
			fmt.Fprintf(&b, "immune to %s", strings.Join(g.immune.words(), ", "))
			if len(g.weak) > 0 {
				fmt.Fprint(&b, "; ")
			}
		}
		if len(g.weak) > 0 {
			fmt.Fprintf(&b, "weak to %s", strings.Join(g.weak.words(), ", "))
		}
		fmt.Fprint(&b, ")")
	}
	fmt.Fprintf(&b, " with an attack that does %d %s damage at initiative %d",
		g.damage, g.damageType, g.init)
	return b.String()
}

func (g *group) inCombat() bool {
	return g.units > 0
}

func (g *group) power() int {
	if g.units < 0 {
		return 0
	}
	return g.units * g.damage
}

func (g *group) damageVs(other *group) int {
	if other.immune[g.damageType] {
		return 0
	}
	dmg := g.power()
	if other.weak[g.damageType] {
		dmg *= 2
	}
	return dmg
}

func (g *group) applyDamage(damage int) int {
	numunits := (damage / g.hp)
	if numunits > g.units {
		numunits = g.units
		g.units = 0
	} else {
		g.units = g.units - numunits
	}
	return numunits
}

func (g *group) bestTarget(allGroups dataType, targetList targets, debug bool) *group {
	valid := make(dataType, 0)
	for _, other := range allGroups {
		if targetList[other] {
			continue
			// don't select already chosen targets
		}
		if other.immune[g.damageType] {
			continue
			// not a valid target if immune
		}
		//if dmg := g.damageVs(other); dmg < other.hp {
		//continue
		//// not a valid target if can't damage
		//}
		if other.army != g.army && other.inCombat() {
			valid = append(valid, other)
		}
	}
	dataBy(func(g1, g2 *group) bool {
		g1dv, g2dv := g.damageVs(g1), g.damageVs(g2)
		if g1dv == g2dv {
			g1pow, g2pow := g1.power(), g2.power()
			if g1pow == g2pow {
				return g1.init > g2.init
			}
			return g1pow > g2pow
		}
		return g1dv > g2dv
	}).Sort(valid)
	if debug {
		for _, v := range valid {
			fmt.Printf("%s group %d would deal defending group %d %d damage\n",
				g.army, g.group, v.group, g.damageVs(v))
		}
	}
	if len(valid) > 0 {
		return valid[0]
	}
	return nil
}

func newGroup(army string, units, hp, damage, init int, damageType, special string) *group {
	g := &group{
		army:       army,
		units:      units,
		hp:         hp,
		damage:     damage,
		init:       init,
		damageType: damageType,
	}
	if special != "" {
		parts := strings.Split(special, "; ")
		for _, p := range parts {
			words := strings.SplitN(p, " ", 3)
			types := strings.Split(words[2], ", ")
			set := make(stringSet)
			for _, t := range types {
				set[t] = true
			}
			if words[0] == "immune" {
				g.immune = set
			} else {
				g.weak = set
			}
		}
	}
	return g
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func readInput(str string) dataType {
	const armyRegex = `^(.*):$`
	const groupRegex = `^(\d+) units each with (\d+) hit points (?:\((.*)\) )?with an attack that does (\d+) (\w+) damage at initiative (\d+)$`

	groups := make(dataType, 0)
	armyRe := regexp.MustCompile(armyRegex)
	groupRe := regexp.MustCompile(groupRegex)
	army := ""
	for _, line := range strings.Split(str, "\n") {
		if line == "" {
			continue
		}
		match := armyRe.FindStringSubmatch(line)
		if len(match) > 0 {
			army = match[1]
			continue
		}
		match = groupRe.FindStringSubmatch(line)
		g := newGroup(army, atoi(match[1]), atoi(match[2]), atoi(match[4]), atoi(match[6]),
			match[5], match[3])
		groups = append(groups, g)
	}
	return groups
}

func solve1(data dataType) int {
	//if 1 == 1 {
	//return 0
	//}
	//	fmt.Print(data)
	army, units := data.runFight(false)
	fmt.Printf("%s won with %d units remaining\n", army, units)
	return units
}

func solve2(data dataType) int {
	//if 1 == 1 {
	//return 0
	//}
	fmt.Print(data)
	for amt := 1; ; amt++ {
		//fmt.Println("===============================================")
		//fmt.Printf("Boost Amount: %d\n", amt)
		data = data.boost(amt)
		//		fmt.Print(data)
		army, units := data.runFight(false)
		fmt.Printf("boost %d - %s won with %d units remaining\n", amt, army, units)
		if army == "Immune System" {
			return units
		}
	}
}
