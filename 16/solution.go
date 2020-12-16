package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type field struct {
	minA, maxA int
	minB, maxB int
	name       string
}

func parseFields(lines []string) []field {
	var result []field
	expression := regexp.MustCompile(`^((?:\w|\s)+)\:\s(\d+)\-(\d+)\sor\s(\d+)\-(\d+)`)
	for _, line := range lines {
		match := expression.FindAllStringSubmatch(line, -1)[0]
		var ints []int
		for _, str := range match[2:] {
			val, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			ints = append(ints, val)
		}
		result = append(result, field{
			name: match[1],
			minA: ints[0],
			maxA: ints[1],
			minB: ints[2],
			maxB: ints[3]})
	}
	return result
}

func intSliceIntersect(sliceA []int, sliceB []int) []int {
	var result []int
	for _, valA := range sliceA {
		for _, valB := range sliceB {
			if valA == valB {
				result = append(result, valA)
			}
		}
	}
	return result
}

func valueFitsField(value int, f field) bool {
	return ((value >= f.minA && value <= f.maxA) || (value >= f.minB && value <= f.maxB))
}

func valueFitsFieldSlice(value int, fieldSlice []field) bool {
	for _, f := range fieldSlice {
		if valueFitsField(value, f) {
			return true
		}
	}
	return false
}

// https://stackoverflow.com/a/19374861
func difference(slice1 []int, slice2 []int) []int {
	var diff []int

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func part1(lines []string) int {
	var fieldEndIndex, nearbyTicketsStartIndex int
	for index, line := range lines {
		if line == "your ticket:" {
			fieldEndIndex = index - 1
			nearbyTicketsStartIndex = index + 4
			break
		}
	}
	fields := parseFields(lines[:fieldEndIndex])
	sum := 0
	for _, line := range lines[nearbyTicketsStartIndex:] {
		values := strings.Split(line, ",")
		for _, value := range values {
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			if !valueFitsFieldSlice(valueInt, fields) {
				sum += valueInt
			}
		}
	}

	return sum
}

func part2(lines []string) int {
	var fieldEndIndex, yourTicketIndex, nearbyTicketsStartIndex int
	for index, line := range lines {
		if line == "your ticket:" {
			fieldEndIndex = index - 1
			yourTicketIndex = index + 1
			nearbyTicketsStartIndex = index + 4
			break
		}
	}

	fields := parseFields(lines[:fieldEndIndex])
	// find all valid values
	var cleanNearbyValues [][]int
	for _, line := range lines[nearbyTicketsStartIndex:] {
		values := strings.Split(line, ",")
		var cleanValueSlice []int
		for _, value := range values {
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			if valueFitsFieldSlice(valueInt, fields) {
				cleanValueSlice = append(cleanValueSlice, valueInt)
			}
		}
		cleanNearbyValues = append(cleanNearbyValues, cleanValueSlice)
	}

	// init possible fields for columns
	var columnPossibleFields [][]int
	for i := 0; i < len(cleanNearbyValues[0]); i++ {
		temp := make([]int, len(fields))
		for j := 0; j < len(fields); j++ {
			temp[j] = j
		}
		columnPossibleFields = append(columnPossibleFields, temp)
	}
	// check possible fields for columns
	for _, ticket := range cleanNearbyValues {
		for columnIndex, column := range ticket {
			var fittingFields []int
			for fieldIndex, f := range fields {
				if valueFitsField(column, f) {
					fittingFields = append(fittingFields, fieldIndex)
				}
			}
			columnPossibleFields[columnIndex] = intSliceIntersect(columnPossibleFields[columnIndex], fittingFields)
		}
	}

	// match columns to fields
	result := make(map[int]string)
	for _, columnA := range columnPossibleFields {
		for index, columnB := range columnPossibleFields {
			if len(columnB) == 1 {
				result[index] = fields[columnB[0]].name
			} else if len(columnB) == len(columnA)+1 {
				result[index] = fields[difference(columnA, columnB)[0]].name
			}
		}
	}

	// calculate final answer
	sum := 1
	yourTicketSplit := strings.Split(lines[yourTicketIndex], ",")
	for columnIndex, name := range result {
		if len(name) >= 9 && name[:9] == "departure" {
			val, err := strconv.Atoi(yourTicketSplit[columnIndex])
			if err != nil {
				log.Fatal(err)
			}
			sum *= val
		}
	}
	return sum
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
