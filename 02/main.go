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

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	catch(err)
	return n
}

type Parsed [][2]int

func parseInput(input string) Parsed {
	parts := strings.Split(input, ",")
	parsed := make(Parsed, len(parts))
	for i, part := range parts {
		parts := strings.Split(part, "-")
		parsed[i] = [2]int{
			atoi(parts[0]),
			atoi(parts[1]),
		}
	}
	return parsed
}

func part1(parsed Parsed) {
	timeStart := time.Now()
	var sum int
	for _, rang := range parsed {
		// brute force the heck out of it
		for val := rang[0]; val <= rang[1]; val++ {
			s := strconv.Itoa(val)
			if len(s)%2 == 0 {
				mid := len(s) / 2
				if s[:mid] == s[mid:] {
					sum += val
				}
			}
		}
	}

	fmt.Printf("Part 1: %d\t\tin %v\n", sum, time.Since(timeStart))
}

func part2(parsed Parsed) {
	timeStart := time.Now()
	var sum int
	for _, rang := range parsed {
		// brute force the heck out of it
		for val := rang[0]; val <= rang[1]; val++ {
			s := strconv.Itoa(val)
			for div := 2; div <= len(s); div++ {
				if len(s)%div == 0 {
					i := len(s) / div
					if strings.Repeat(s[:i], div) == s {
						sum += val
						break
					}
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\t\tin %v\n", sum, time.Since(timeStart))
}
