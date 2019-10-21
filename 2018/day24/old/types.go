package main

import (
	"fmt"
	"strings"
)

type StringSet map[string]bool

type Group struct {
	Army       string
	Units      int
	HP         int
	immune     StringSet
	weak       StringSet
	Damage     int
	DamageType string
	Initiative int
	Target     *Group
}

func NewGroup(army string, units, hp, damage, initiative int, damageType, special string) *Group {
	g := Group{}
	g.Army = army
	g.Units = units
	g.HP = hp
	g.Damage = damage
	g.Initiative = initiative
	g.DamageType = damageType
	if special != "" {
		parts := strings.Split(special, "; ")
		for _, p := range parts {
			words := strings.SplitN(p, " ", 3)
			types := strings.Split(words[2], ", ")
			set := make(StringSet)
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
	return &g
}

func (g *Group) EffectivePower() int {
	return g.Units * g.Damage
}

func (g *Group) PowerVs(target *Group) int {
	if target.Army == g.Army {
		return 0
	}
	if target.immune[g.DamageType] {
		return 0
	}
	dmg := g.EffectivePower()
	if target.weak[g.DamageType] {
		return 2 * dmg
	}
	return dmg
}

func (g *Group) String() string {
	var out strings.Builder
	fmt.Fprintf(&out, "<%s %d@%d (%d %s) on %d", g.Army, g.Units, g.HP, g.Damage, g.DamageType, g.Initiative)
	if len(g.immune) > 0 {
		immunes := make([]string, 0, len(g.immune))
		for i := range g.immune {
			immunes = append(immunes, i)
		}
		fmt.Fprintf(&out, " im: %s", strings.Join(immunes, ", "))
	}
	if len(g.weak) > 0 {
		weaks := make([]string, 0, len(g.weak))
		for i := range g.weak {
			weaks = append(weaks, i)
		}
		fmt.Fprintf(&out, " wk: %s", strings.Join(weaks, ", "))
	}
	out.WriteString(">")
	return out.String()
}
