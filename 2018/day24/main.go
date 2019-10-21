package main

import "container/list"
import "fmt"
import "sort"

func stillFighting(groups []*Group) bool {
	counts := make(map[string]int)
	for _, g := range groups {
		if g.Units > 0 {
			counts[g.Army] = counts[g.Army] + 1
		}
	}
	return len(counts) > 1
}

func targetSelection(groups []*Group) []*Group {
	fighters := make([]*Group, 0, len(groups))
	candidates := list.New()
	for _, g := range groups {
		if g.Units > 0 {
			fighters = append(fighters, g)
			candidates.PushBack(g)
		}
	}
	sort.Slice(fighters, func(i, j int) bool {
		ip := fighters[i].EffectivePower()
		jp := fighters[j].EffectivePower()
		if ip == jp {
			return fighters[i].Initiative > fighters[j].Initiative
		}
		return ip > jp
	})
	for _, f := range fighters {
		var targetE *list.Element
		var targetDamage int
		var targetPower int
		var targetInit int
		for e := candidates.Front(); e != nil; e = e.Next() {
			t := e.Value.(*Group)
			if t.Army == f.Army {
				continue
			}
			damage := f.PowerVs(t)
			if damage > targetDamage {
				targetDamage = damage
				targetPower = t.EffectivePower()
				targetInit = t.Initiative
				targetE = e
			} else if damage == targetDamage {
				power := t.EffectivePower()
				if power > targetPower {
					targetPower = power
					targetInit = t.Initiative
					targetE = e
				} else if power == targetPower {
					if t.Initiative > targetInit {
						targetInit = t.Initiative
						targetE = e
					}
				}
			}
		}
		if targetE == nil {
			f.Target = nil
		} else {
			f.Target = targetE.Value.(*Group)
			candidates.Remove(targetE)
		}
	}
	return fighters
}

func attack(fighters []*Group) {
	sort.Slice(fighters, func(i, j int) bool {
		return fighters[i].Initiative > fighters[j].Initiative
	})
	for _, f := range fighters {
		if f.Target == nil {
			continue
		}
		if f.Units <= 0 {
			continue
		}
		totalAttack := f.PowerVs(f.Target)
		totalUnits := totalAttack / f.Target.HP
		f.Target.Units -= totalUnits
		if f.Target.Units < 0 {
			f.Target.Units = 0
		}
		//fmt.Println(f, "(VS)", f.Target)
		//fmt.Println("    ", totalAttack, "destroying", totalUnits, "leaving", f.Target.Units)
	}
}

func main() {
	part1()
	part2()
}

func runFight(str string, boost int) (army string, units int) {
	allGroups := parseInput(str)
	for i := range allGroups {
		if allGroups[i].Army != "Infection" {
			allGroups[i].Damage += boost
		}
	}
	for stillFighting(allGroups) {
		fighters := targetSelection(allGroups)
		attack(fighters)
	}
	army = "Infection"
	units = 0
	for _, g := range allGroups {
		if g.Units > 0 {
			army = g.Army
			units += g.Units
		}
	}
	return
}

func part1() {
	str := inputStr
	//str = testStr
	army, total := runFight(str, 0)
	fmt.Println("Part 1:", army, total)
}

func part2() {
	str := inputStr
	var cur, lo, hi = 0, 0, 10001
	var army string
	var total int
	for lo < hi {
		army, total = runFight(str, cur)
		fmt.Printf("  Tried a boost of %d (range %d - %d): %s %d\n", cur, lo, hi, army, total)
		if army == "Infection" {
			lo = cur + 1
		} else {
			hi = cur - 1
		}
		cur = (lo + hi) / 2
	}
	army, total = runFight(str, cur)
	fmt.Printf("  Tried a boost of %d (range %d - %d): %s %d\n", cur, lo, hi, army, total)
	fmt.Printf("  ------------------")
	prev := 0
	for boost := 0; boost < 42; boost++ {
		army, total = runFight(str, boost)
		fmt.Printf("  Tried a boost of %d: %s %d (Δ%d)\n", boost, army, total, prev-total)
		prev = total
	}

	fmt.Println("Part 2", army, total, cur)
}
