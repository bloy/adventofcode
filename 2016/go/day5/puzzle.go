package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const inputText = "ugkcyxxp"
const pwdSize = 8

func readInput(inputStr string) string {
	return inputStr
}

func solve1(roomCode string) string {
	var pos, index int
	pwd := make([]rune, pwdSize)
	for i := 0; i < pwdSize; i++ {
		pwd[i] = '_'
	}
	fmt.Print("    Running Part 1: ________")
	for pos < pwdSize {
		hash := "abcdefgh"
		for hash[:5] != "00000" {
			str := fmt.Sprintf("%s%d", roomCode, index)
			hash = fmt.Sprintf("%x", md5.Sum([]byte(str)))
			index++
		}
		pwd[pos] = rune(hash[5])
		fmt.Print("")
		fmt.Print(string(pwd))
		pos++
	}
	fmt.Println("")
	return string(pwd)
}

func solve2(roomCode string) string {
	pwd := make([]rune, pwdSize)
	for i := 0; i < pwdSize; i++ {
		pwd[i] = '_'
	}
	var index int
	fmt.Print("    Running Part 2: ________")
	for !pwdComplete(pwd) {
		hash := "abcdefgh"
		for hash[:5] != "00000" || rune(hash[5]) < '0' || rune(hash[5]) > '7' {
			str := fmt.Sprintf("%s%d", roomCode, index)
			hash = fmt.Sprintf("%x", md5.Sum([]byte(str)))
			index++
		}
		pos, err := strconv.Atoi(string(hash[5]))
		if err != nil {
			panic(err)
		}
		if pwd[pos] != '_' {
			continue
		}
		pwd[pos] = rune(hash[6])
		fmt.Print("")
		fmt.Print(string(pwd))
	}
	fmt.Println()
	return string(pwd)
}

func pwdComplete(pwd []rune) bool {
	for i := 0; i < len(pwd); i++ {
		if pwd[i] == '_' {
			return false
		}
	}
	return true
}
