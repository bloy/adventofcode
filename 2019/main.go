package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)
	probNum, solNum, err := parseArguments()
	if err != nil {
		log.Fatalln("Error while parsing arguments", err)
	}
	soln, err := FindSolution(probNum, solNum)
	if err != nil {
		log.Fatalln(err)
	}
	pr, err := newPuzzleRun(probNum)
	if err != nil {
		log.Fatalln(err)
	}
	defer pr.Close()

	pr.timings.start = time.Now()
	for _, part := range soln.Parts {
		part(pr)
	}
	pr.timings.done = time.Now()
	pr.printTimings()
}

func parseArguments() (probNum, solNum int, err error) {
	flag.Parse()
	if flag.NArg() != 1 {
		err = fmt.Errorf("No problem specified as arguments")
		return
	}
	prbStr := flag.Arg(0)
	parts := strings.SplitN(prbStr, ".", 2)
	solNum = 1
	if len(parts) == 2 {
		solNum, err = strconv.Atoi(parts[1])
		if err != nil {
			return
		}
	}
	probNum, err = strconv.Atoi(parts[0])
	if err != nil {
		return
	}
	return
}
