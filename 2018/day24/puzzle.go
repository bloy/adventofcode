package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type dataType []*group

func (data dataType) String() string {
	b := strings.Builder{}
	c := make(dataType, len(data))
	copy(c, data)
	dataBy(func(g1, g2 *group) bool { return g1.army < g2.army }).Sort(c)
	army := ""
	for i := range c {
		if c[i].army != army {
			fmt.Fprintf(&b, "\n%s:\n", c[i].army)
			army = c[i].army
		}
		if !c[i].inCombat() {
			continue
		}
		fmt.Fprintf(&b, "%v\n", c[i])
	}
	return b.String()
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
	units, hp, damage, init int
	immune, weak            stringSet
	army, damageType        string
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
	fmt.Print(data)
	return 0
}

func solve2(data dataType) int {
	return 0
}
