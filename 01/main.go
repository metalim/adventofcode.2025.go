package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: go run . input.txt")
		os.Exit(1)
	}

	bs, err := os.ReadFile(flag.Arg(0))
	catch(err)

	parsed := parseInput(string(bs))
	part1(parsed)
	part2(parsed)
}

type Parsed []string

func parseInput(input string) Parsed {
	lines := strings.Split(input, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	return Parsed(lines)
}

func part1(parsed Parsed) {
	timeStart := time.Now()
	dial := 50
	var zeroes int
	for _, line := range parsed {
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		catch(err)
		switch dir {
		case 'L':
			dial -= dist
		case 'R':
			dial += dist
		default:
			panic(fmt.Sprintf("invalid direction: %c", dir))
		}
		dial = (dial + 100) % 100
		if dial == 0 {
			zeroes++
		}
	}

	fmt.Printf("Part 1: %d\t\tin %v\n", zeroes, time.Since(timeStart))
}

func part2(parsed Parsed) {
	timeStart := time.Now()
	dial := 50
	var zeroes int
	for _, line := range parsed {
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		catch(err)
		var delta int
		switch dir {
		case 'L':
			delta = -1
		case 'R':
			delta = 1
		default:
			panic(fmt.Sprintf("invalid direction: %c", dir))
		}
		for ; dist > 0; dist-- {
			dial = (dial + delta + 100) % 100
			if dial == 0 {
				zeroes++
			}
		}
	}

	fmt.Printf("Part 2: %d\t\tin %v\n", zeroes, time.Since(timeStart))
}
