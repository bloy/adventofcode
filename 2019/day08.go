package main

import (
	"bufio"
	"fmt"
	"strings"
)

func init() {
	AddSolution(8, solveDay8)
}

func solveDay8(pr *PuzzleRun) {
	scanner := bufio.NewScanner(pr.InFile)
	var imagedata string
	var xsize, ysize int
	for scanner.Scan() {
		imagedata += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		pr.logger.Fatal(err)
	}
	xsize, ysize = 25, 6
	pr.ReportLoad()

	datalen := len(imagedata)
	numlayers := datalen / xsize / ysize
	layers := make([]string, numlayers)
	min := xsize*ysize + 1
	minLayer := numlayers + 1
	for i := 0; i < numlayers; i++ {
		layers[i] = imagedata[i*xsize*ysize : (i+1)*xsize*ysize]
		c := strings.Count(layers[i], "0")
		if c < min {
			min = c
			minLayer = i
		}
	}
	count1 := strings.Count(layers[minLayer], "1")
	count2 := strings.Count(layers[minLayer], "2")
	pr.ReportPart(count1 * count2)

	img := make([]rune, xsize*ysize)
	for i := 0; i < xsize*ysize; i++ {
		img[i] = '.'
	}
	buf := strings.Builder{}
	for y := 0; y < ysize; y++ {
		fmt.Fprint(&buf, "\n")
		for x := 0; x < xsize; x++ {
			c := "."
			for i := 0; i < numlayers; i++ {
				layer := layers[numlayers-i-1]
				switch layer[y*xsize+x] {
				case '0':
					c = " "
				case '1':
					c = "#"
				default:
				}
			}
			fmt.Fprint(&buf, c)
		}
	}
	fmt.Fprint(&buf, "\n")
	pr.ReportPart(buf.String())
}
