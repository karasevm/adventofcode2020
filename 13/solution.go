package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"regexp"
	"strconv"
)

func part1(lines []string) int {
	digitExpression := regexp.MustCompile(`\d{1,4}`)
	intervalsStrings := digitExpression.FindAllString(lines[1], -1)
	var intervals []int
	for _, s := range intervalsStrings {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		intervals = append(intervals, num)
	}
	startTime, err := strconv.Atoi(lines[0])
	if err != nil {
		log.Fatal(err)
	}
	minDiff := math.MaxInt32
	minDiffInterval := 0
	for _, interval := range intervals {
		multiplier := 1
		for interval*multiplier < startTime {
			multiplier++
		}
		currentDiff := (interval * multiplier) - startTime
		if currentDiff < minDiff {
			minDiff = currentDiff
			minDiffInterval = interval
		}

	}
	return minDiff * minDiffInterval
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b uint64) uint64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b uint64) uint64 {
	var bigA, bigB, bigGCD big.Int
	bigA.SetUint64(a)
	bigB.SetUint64(b)
	bigGCD.SetUint64(GCD(a, b))
	bigA.Mul(&bigA, &bigB)
	bigA.Div(&bigA, &bigGCD)
	return bigA.Uint64()
}

func part2(lines []string) uint64 {
	digitExpression := regexp.MustCompile(`\d{1,4}|x`)
	intervalsStrings := digitExpression.FindAllString(lines[1], -1)
	var intervals []int
	for _, s := range intervalsStrings {
		if s == "x" {
			intervals = append(intervals, 1)
		} else {

			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			intervals = append(intervals, num)
		}
	}

	var nextMultipliers []uint64
	var multiplier uint64
	multiplierOffset := uint64(1)
	multiplierStep := uint64(1)
	for index, interval := range intervals[1:] {
		nextMultipliers = nil
		for multiplier = multiplierOffset; len(nextMultipliers) < 2; multiplier += multiplierStep {
			multiple := uint64(intervals[index]) * multiplier
			if (multiple+1)%uint64(interval) == 0 {
				nextMultipliers = append(nextMultipliers, (multiple+1)/uint64(interval))
			}
		}
		multiplierStep = LCM(multiplierStep, (nextMultipliers[1] - nextMultipliers[0]))
		multiplierOffset = nextMultipliers[0]
	}
	return nextMultipliers[0]*uint64(intervals[len(intervals)-1]) - uint64(len(intervals)-1)
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		log.Fatal("Provide the path to input")
	}
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	fmt.Printf("Part 1 answer:%v\nPart 2 answer:%v\n", part1(lines), part2(lines))
}
