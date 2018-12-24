package main

import "container/list"
import "fmt"
import "sort"

func stillFighting(groups []*Group) bool {
	army := ""
	for _, g := range groups {
		if g.Units > 0 && army != "" && g.Army != army {
			return true
		}
		army = g.Army
	}
	return false
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
		if targetE != nil {
			f.Target = targetE.Value.(*Group)
			candidates.Remove(targetE)
		}
	}
	return fighters
}

func main() {
	str := inputStr
	str = testStr
	allGroups := parseInput(str)
	for _, g := range allGroups {
		fmt.Println(g)
	}

	//for stillFighting(allGroups) {
	fighters := targetSelection(allGroups)
	for _, f := range fighters {
		fmt.Println(f, "vs", f.Target)
	}
	//}
}
