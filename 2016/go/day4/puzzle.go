package main

import (
	"sort"
	"strconv"
	"strings"
)

// Rooms holds a list of Room
type Rooms []*Room

// Room is a room object
type Room struct {
	Name     string
	SectorID int
	Checksum string
}

// NewRoom creates a new room
func NewRoom(name string, sector int, checksum string) *Room {
	return &Room{Name: name, SectorID: sector, Checksum: checksum}
}

// IsReal checks if a room is a real room
func (r *Room) IsReal() bool {
	letterCount := make(map[rune]int)
	for _, letter := range r.Name {
		if letter == '-' {
			continue
		}
		letterCount[letter] = letterCount[letter] + 1
	}
	letters := make([]rune, 0, len(letterCount))
	for k := range letterCount {
		letters = append(letters, k)
	}
	sort.Slice(letters, func(i, j int) bool {
		iCount := letterCount[letters[i]]
		jCount := letterCount[letters[j]]
		if iCount == jCount {
			return letters[i] < letters[j]
		}
		return iCount > jCount
	})
	return string(letters[:5]) == r.Checksum
}

// DecryptedName decrypts the room's name
func (r *Room) DecryptedName() string {
	chars := make([]rune, len(r.Name))
	for i, c := range r.Name {
		if c == '-' {
			chars[i] = ' '
			continue
		}
		d := ((c - 'a' + rune(r.SectorID)) % 26) + 'a'
		chars[i] = d
	}
	return string(chars)
}

func readInput(inputText string) Rooms {
	lines := strings.Split(inputText, "\n")
	rooms := make(Rooms, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, "-")
		name := strings.Join(parts[:len(parts)-1], "-")
		parts = strings.Split(parts[len(parts)-1], "[")
		sector, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		checksum := parts[1][:len(parts[1])-1]
		rooms[i] = NewRoom(name, sector, checksum)
	}
	return rooms
}

func solve1(rooms Rooms) int {
	sum := 0
	for _, r := range rooms {
		if r.IsReal() {
			sum += r.SectorID
		}
	}
	return sum
}

func solve2(rooms Rooms) int {
	for _, r := range rooms {
		if r.IsReal() {
			if r.DecryptedName() == "northpole object storage" {
				return r.SectorID
			}
		}
	}
	return 0
}
