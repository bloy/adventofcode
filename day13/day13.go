package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const testStr string = `/->-\
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
`

var turnDir map[string]map[string]string = map[string]map[string]string{
	"L": {
		"<": "v",
		"v": ">",
		">": "^",
		"^": "<",
	},
	"R": {
		"<": "^",
		"^": ">",
		">": "v",
		"v": "<",
	},
	"S": {
		"<": "<",
		">": ">",
		"v": "v",
		"^": "^",
	},
}

var nextTurn map[string]string = map[string]string{
	"L": "S",
	"S": "R",
	"R": "L",
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("<%d,%d>", p.x, p.y)
}

type Cart struct {
	direction string
	nextTurn  string
	start     Point
	position  Point
}

func (c Cart) String() string {
	return fmt.Sprintf("[%d: %s(%s) %v]", c.direction, c.nextTurn, c.position)
}

type Carts []Cart

func (carts Carts) Len() int      { return len(carts) }
func (carts Carts) Swap(i, j int) { carts[i], carts[j] = carts[j], carts[i] }
func (carts Carts) Less(i, j int) bool {
	if carts[i].position.y == carts[j].position.y {
		return carts[i].position.x < carts[j].position.x
	}
	return carts[i].position.y < carts[j].position.y
}

type Tracks [][]rune

func getInput() (tracks Tracks, carts Carts) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	contentStr := string(content)
	//contentStr = testStr
	carts = make(Carts, 0)
	tracks = make(Tracks, 0)
	for y, str := range strings.Split(contentStr, "\n") {
		if str == "" {
			continue
		}
		tracks = append(tracks, make([]rune, 0, len(str)))
		for x, ch := range str {
			track := ch
			switch ch {
			case '^':
				carts = append(carts, Cart{"^", "L", Point{x, y}, Point{x, y}})
				track = '|'
			case 'v':
				carts = append(carts, Cart{"v", "L", Point{x, y}, Point{x, y}})
				track = '|'
			case '<':
				carts = append(carts, Cart{"<", "L", Point{x, y}, Point{x, y}})
				track = '-'
			case '>':
				carts = append(carts, Cart{">", "L", Point{x, y}, Point{x, y}})
				track = '-'
			}
			tracks[y] = append(tracks[y], track)
		}
	}
	sort.Sort(carts)
	return
}

func printTracks(tracks Tracks) {
	maxLen := 0
	for _, row := range tracks {
		if maxLen < len(row) {
			maxLen = len(row)
		}
	}
	fmt.Print("   ")
	for i := 0; i < maxLen; i++ {
		if i >= 100 {
			fmt.Print(i / 100 % 10)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n   ")
	for i := 0; i < maxLen; i++ {
		if i >= 10 {
			fmt.Print(i / 10 % 10)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n   ")
	for i := 0; i < maxLen; i++ {
		fmt.Print(i % 10)
	}
	fmt.Print("\n")
	for y, row := range tracks {
		fmt.Printf("%3d", y)
		for _, ch := range row {
			fmt.Printf("%s", string(ch))
		}
		fmt.Print("\n")
	}
}

func printCarts(carts Carts) {
	for _, cart := range carts {
		fmt.Printf("[%s,%3d,%3d] ", cart.direction, cart.position.x, cart.position.y)
	}
	fmt.Print("\n")
}

func stepSim(tracks *Tracks, carts *Carts) {
	sort.Sort(carts)
	for i := 0; i < len(*carts); i++ {
		cart := (*carts)[i]
		switch cart.direction {
		case "^":
			cart.position.y--
		case "v":
			cart.position.y++
		case ">":
			cart.position.x++
		case "<":
			cart.position.x--
		}
		track := (*tracks)[cart.position.y][cart.position.x]
		switch track {
		case '\\':
			switch cart.direction {
			case "^":
				cart.direction = "<"
			case "<":
				cart.direction = "^"
			case "v":
				cart.direction = ">"
			case ">":
				cart.direction = "v"
			}
		case '/':
			switch cart.direction {
			case "^":
				cart.direction = ">"
			case ">":
				cart.direction = "^"
			case "v":
				cart.direction = "<"
			case "<":
				cart.direction = "v"
			}
		case '+':
			cart.direction = turnDir[cart.nextTurn][cart.direction]
			cart.nextTurn = nextTurn[cart.nextTurn]
		}
		(*carts)[i] = cart
		for j := 0; j < len(*carts); j++ {
			if i == j {
				continue
			}
			if cart.position == (*carts)[j].position {
				return
			}
		}
	}
}

func checkCrashed(cartsPtr *Carts) (bool, *Point) {
	carts := *cartsPtr
	for i := 0; i < len(carts)-1; i++ {
		for j := i + 1; j < len(carts); j++ {
			if carts[i].position == carts[j].position {
				return true, &Point{carts[i].position.x, carts[i].position.y}
			}
		}
	}
	return false, nil
}

func runPart1(tracks Tracks, carts Carts) {
	printTracks(tracks)
	printCarts(carts)
	var crashed bool = false
	var crashPoint *Point
	for !crashed {
		stepSim(&tracks, &carts)
		printCarts(carts)
		crashed, crashPoint = checkCrashed(&carts)
	}
	fmt.Println("Part 1", crashPoint)
}

func main() {
	tracks, carts := getInput()

	runPart1(tracks, carts)
}
