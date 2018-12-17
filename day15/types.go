package main

import (
	"container/list"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Point struct {
	x, y int
}

type Unit struct {
	Power int
	HP    int
	Elf   bool
	Dead  bool
}

func (u Unit) String() string {
	c := "G"
	if u.Elf {
		c = "E"
	}
	return fmt.Sprintf("%s(%3d)", c, u.HP)
}

type SquareType byte

const (
	Space SquareType = iota
	Wall
)

type Square struct {
	Unit  *Unit
	Kind  SquareType
	Pos   Point
	Level *Level
}

func (s Square) String() string {
	switch s.Kind {
	case Space:
		if s.Unit == nil {
			return "."
		} else {
			if s.Unit.Elf {
				return "E"
			} else {
				return "G"
			}
		}
	case Wall:
		return "#"
	default:
		return " "
	}
}

type Level struct {
	Xsize   int
	Ysize   int
	squares []Square
	rounds  int
}

func (level *Level) UnitSquares() []*Square {
	var squares []*Square = make([]*Square, 0)
	for i := 0; i < len(level.squares); i++ {
		square := level.squares[i]
		if square.Unit != nil && !square.Unit.Dead {
			squares = append(squares, &square)
		}
	}
	return squares
}

func (level *Level) AdjacentSquares(s Square) []*Square {
	adjacent := make([]*Square, 0, 4)
	adjacent = append(adjacent, &(level.squares[(s.Pos.y-1)*level.Xsize+s.Pos.x]))
	adjacent = append(adjacent, &(level.squares[s.Pos.y*level.Xsize+s.Pos.x-1]))
	adjacent = append(adjacent, &(level.squares[s.Pos.y*level.Xsize+s.Pos.x+1]))
	adjacent = append(adjacent, &(level.squares[(s.Pos.y+1)*level.Xsize+s.Pos.x]))
	return adjacent
}

func (level *Level) TargetSquares(s Square) []*Square {
	if s.Unit == nil {
		return []*Square{}
	}
	targets := make([]*Square, 0)
	for i := 0; i < len(level.squares); i++ {
		square := level.squares[i]
		if square.Unit != nil && square.Unit.Elf != s.Unit.Elf && !square.Unit.Dead {
			for _, try := range level.AdjacentSquares(square) {
				if *try == s {
					return []*Square{try}
				}
				if try.Kind != Wall && (try.Unit == nil) {
					targets = append(targets, try)
				}
			}
		}
	}
	sort.Slice(targets, func(i, j int) bool {
		if targets[i].Pos.y == targets[j].Pos.y {
			return targets[i].Pos.x < targets[j].Pos.x
		}
		return targets[i].Pos.y < targets[j].Pos.y
	})
	return targets
}

type SquareStep struct {
	square *Square
	steps  int
}

func (level *Level) PathLength(s1, s2 Square) (int, error) {
	if s1 == s2 {
		return 0, nil
	}
	stepQueue := list.New()
	seen := make(map[*Square]bool)
	seen[&s1] = true
	for _, s := range level.AdjacentSquares(s1) {
		if s.Kind != Wall && s.Unit == nil {
			stepQueue.PushBack(SquareStep{s, 1})
		}
	}
	for stepQueue.Front() != nil {
		step := stepQueue.Front().Value.(SquareStep)
		stepQueue.Remove(stepQueue.Front())
		if *(step.square) == s2 {
			return step.steps, nil
		}
		for _, s := range level.AdjacentSquares(*(step.square)) {
			if _, ok := seen[s]; !ok && s.Kind != Wall && s.Unit == nil {
				stepQueue.PushBack(SquareStep{s, 1 + step.steps})
				seen[s] = true
			}
		}
	}
	return 0, errors.New("No path to target square")
}

func (level *Level) HasTargets(s *Square) bool {
	if s.Unit == nil {
		return true
	}
	for i := 0; i < len(level.squares); i++ {
		square := level.squares[i]
		if *s != square && square.Unit != nil && square.Unit.Elf != s.Unit.Elf && !square.Unit.Dead {
			return true
		}
	}
	return false
}

func (level *Level) MoveAction(s *Square) *Square {
	// perform the move action for the unit in square s
	if s.Unit == nil || s.Unit.Dead {
		return s
	}
	pathLength := 999
	var target *Square
	for _, t := range level.TargetSquares(*s) {
		l, err := level.PathLength(*s, *t)
		if err == nil && l < pathLength {
			pathLength = l
			target = t
		}
	}
	if pathLength == 0 && *target == *s {
		return s
	}
	if target == nil {
		return s
	}
	var step *Square
	for _, t := range level.AdjacentSquares(*s) {
		if t.Kind == Wall || t.Unit != nil {
			continue
		}
		l, err := level.PathLength(*t, *target)
		if err == nil && l < pathLength {
			pathLength = l
			step = t
		}
	}
	if step == nil {
		panic(fmt.Sprintf("oh no s = %#v\n", s))
	}
	step.Unit = s.Unit
	s.Unit = nil
	return step
}

func (level *Level) AttackAction(s *Square) bool {
	// returns true iff an elf dies as a result of this action
	if s.Unit == nil || s.Unit.Dead {
		return false
	}
	var targetSquare *Square
	var minHP int = 500
	for _, t := range level.AdjacentSquares(*s) {
		if t.Unit != nil && t.Unit.Elf != s.Unit.Elf && t.Unit.HP < minHP {
			minHP = t.Unit.HP
			targetSquare = t
		}
	}
	if targetSquare == nil {
		return false
	}
	targetSquare.Unit.HP -= s.Unit.Power
	if targetSquare.Unit.HP <= 0 {
		targetSquare.Unit.Dead = true
		unit := targetSquare.Unit
		targetSquare.Unit = nil
		if unit.Elf {
			return true
		}
	}
	return false
}

func (level *Level) PlayRound() (fullRound bool, elfDeath bool) {
	fullRound = true
	elfDeath = false
	unitSquares := list.New()
	for i := 0; i < len(level.squares); i++ {
		if level.squares[i].Unit != nil {
			unitSquares.PushBack(&(level.squares[i]))
		}
	}
	for unitSquares.Front() != nil {
		s := unitSquares.Front().Value.(*Square)
		if !level.HasTargets(s) {
			fullRound = false
			return
		}
		unitSquares.Remove(unitSquares.Front())
		s = level.MoveAction(s)
		elfDeath = elfDeath || level.AttackAction(s)
	}
	level.rounds++
	return
}

func NewLevel(str string, power int) (level *Level) {
	str = strings.TrimSpace(str)
	lines := strings.Split(str, "\n")
	level = &Level{}
	level.rounds = 0
	level.Ysize = len(lines)
	level.Xsize = len(lines[0])
	level.squares = make([]Square, level.Xsize*level.Ysize)
	for y, line := range lines {
		for x, ch := range line {
			square := Square{}
			square.Level = level
			square.Kind = Space
			square.Pos = Point{x, y}
			switch ch {
			case '#':
				square.Kind = Wall
			case 'G':
				unit := &Unit{3, 200, false, false}
				square.Unit = unit
			case 'E':
				unit := &Unit{power, 200, true, false}
				square.Unit = unit
			default:
				square.Kind = Space
			}
			level.squares[y*level.Xsize+x] = square
		}
	}
	return level
}

func (level *Level) String() string {
	var out strings.Builder
	unitList := list.New()
	fmt.Fprintf(&out, "After %d rounds:\n", level.rounds)
	for y := 0; y < level.Ysize; y++ {
		for x := 0; x < level.Xsize; x++ {
			square := level.squares[y*level.Xsize+x]
			fmt.Fprintf(&out, "%v", square)
			if square.Unit != nil {
				unitList.PushBack(square.Unit)
			}
		}
		fmt.Fprint(&out, "  ")
		for unitList.Front() != nil {
			fmt.Fprintf(&out, " %v", unitList.Front().Value)
			unitList.Remove(unitList.Front())
		}
		fmt.Fprint(&out, "\n")
	}
	return out.String()
}

func (level *Level) Score() string {
	var out strings.Builder
	var winner string
	var hptotal int
	for _, s := range level.squares {
		if s.Unit != nil {
			hptotal += s.Unit.HP
			if s.Unit.Elf {
				winner = "Elves"
			} else {
				winner = "Goblins"
			}
		}
	}
	fmt.Fprintf(&out, "Combat ends after %d full rounds\n", level.rounds)
	fmt.Fprintf(&out, "%s win with %d total hit points left\n", winner, hptotal)
	fmt.Fprintf(&out, "%d * %d = %d", level.rounds, hptotal, level.rounds*hptotal)
	return out.String()
}
