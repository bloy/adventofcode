package main

import "fmt"

func doPart1() {
	bots := ParseStr(inputStr)
	maxBot := bots[0]
	for i := range bots {
		if bots[i].Radius > maxBot.Radius {
			maxBot = bots[i]
		}
	}
	botCount := 0
	for i := range bots {
		if bots[i].Position.Distance(maxBot.Position) <= maxBot.Radius {
			botCount++
		}
	}

	fmt.Println("Part 1 bot count:", botCount)
}

func doPart2() {
	bots := ParseStr(inputStr)
	minPoint := bots[0].Position
	maxPoint := bots[0].Position
	for _, bot := range bots {
		if bot.Position.X < minPoint.X {
			minPoint.X = bot.Position.X
		}
		if bot.Position.X > maxPoint.X {
			maxPoint.X = bot.Position.X
		}
		if bot.Position.Y < minPoint.Y {
			minPoint.Y = bot.Position.Y
		}
		if bot.Position.Y > maxPoint.Y {
			maxPoint.Y = bot.Position.Y
		}
		if bot.Position.Z < minPoint.Z {
			minPoint.Z = bot.Position.Z
		}
		if bot.Position.Z > maxPoint.Z {
			maxPoint.Z = bot.Position.Z
		}
	}

	boxSize := 1
	for boxSize < maxPoint.X-minPoint.X {
		boxSize <<= 1 // set the initial search box size
	}
	boxSize <<= 1

	zeroPoint := Point{0, 0, 0}
	for boxSize != 1 {
		boxSize >>= 1
		bestCount := 0
		bestPoint := Point{}
		bestDistance := -1
		for x := minPoint.X; x <= maxPoint.X; x += boxSize {
			for y := minPoint.Y; y <= maxPoint.Y; y += boxSize {
				for z := minPoint.Z; z <= maxPoint.Z; z += boxSize {
					count := 0
					for _, bot := range bots {
						p1 := Point{bot.Position.X, bot.Position.Y, bot.Position.Z}
						p2 := Point{x, y, z}
						r := bot.Radius
						if boxSize > 1 {
							p1.X = p1.X / boxSize
							p1.Y = p1.Y / boxSize
							p1.Z = p1.Z / boxSize
							p2.X = p2.X / boxSize
							p2.Y = p2.Y / boxSize
							p2.Z = p2.Z / boxSize
							r = r/boxSize + 1 // +1 for edge effects
						}
						if p1.Distance(p2) <= r {
							count++
						}
					}
					if count > bestCount {
						bestCount = count
						bestPoint = Point{x, y, z}
						bestDistance = bestPoint.Distance(zeroPoint)
					} else if count == bestCount {
						if bestDistance == -1 || zeroPoint.Distance(Point{x, y, z}) < bestDistance {
							bestPoint = Point{x, y, z}
							bestDistance = zeroPoint.Distance(bestPoint)
						}
					}
				}
			}
		}
		if boxSize == 1 {
			fmt.Println("Part 2: max count =", bestCount, "distance = ", bestDistance)
			break
		}
		// still searching and reducing box size
		minPoint = Point{bestPoint.X - boxSize, bestPoint.Y - boxSize, bestPoint.Z - boxSize}
		maxPoint = Point{bestPoint.X + boxSize, bestPoint.Y + boxSize, bestPoint.Z + boxSize}
	}
}

func main() {
	doPart1()
	doPart2()
}
