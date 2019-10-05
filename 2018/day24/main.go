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
		var targetE *list.Element = nil
		var targetDamage int = 0
		var targetPower int = 0
		var targetInit int = 0
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
	str := inputStr
	//str = testStr
	allGroups := parseInput(str)

	for stillFighting(allGroups) {
		fighters := targetSelection(allGroups)
		attack(fighters)
	}
	total := 0
	for _, g := range allGroups {
		if g.Units > 0 {
			total += g.Units
		}
	}
	fmt.Println("Part 1:", total)
}
