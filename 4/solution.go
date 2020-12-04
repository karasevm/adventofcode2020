package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func cropResult(matchResult [][]string) [][]string {
	var newSlice [][]string
	for _, slice := range matchResult {
		newSlice = append(newSlice, []string{slice[1], slice[2]})
	}
	return newSlice
}

func parsePassports(lines []string) [][][]string {
	var passports [][][]string
	expression := regexp.MustCompile(`(\w{3}):([^\s]+)`)
	var passport [][]string
	for _, line := range lines {
		if len(line) == 0 {
			passports = append(passports, passport)
			passport = nil
			continue
		}
		match := expression.FindAllStringSubmatch(line, -1)
		if len(match) > 0 {
			passport = append(passport, cropResult(match)...)
		}
	}
	return passports
}

func findValidPassports(lines []string) [][][]string {
	var validPassports [][][]string
	for _, passport := range parsePassports(lines) {
		fieldCount := len(passport)
		for _, field := range passport {
			if field[0] == "cid" {
				fieldCount--
				break
			}
		}
		if fieldCount >= 7 {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}

func part1(lines []string) int {
	return len(findValidPassports(lines))
}

func verifyBirthYear(year string) bool {
	value, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return (value >= 1920) && (value <= 2002)
}

func verifyIssueYear(year string) bool {
	value, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return (value >= 2010) && (value <= 2020)
}

func verifyExpYear(year string) bool {
	value, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return (value >= 2020) && (value <= 2030)
}

func verifyHeight(height string) bool {
	value, err := strconv.Atoi(height[:len(height)-2])
	if err != nil {
		return false
	}
	unit := height[len(height)-2:]
	switch unit {
	case "cm":
		return (value >= 150) && (value <= 193)
	case "in":
		return (value >= 59) && (value <= 76)
	default:
		return false
	}
}

func verifyHairColor(color string) bool {
	hexExpression := regexp.MustCompile(`\#[[:xdigit:]]{6}`)
	return len(hexExpression.FindAllString(color, -1)) != 0
}

func verifyEyeColor(color string) bool {
	switch color {
	case
		"amb",
		"blu",
		"brn",
		"grn",
		"gry",
		"hzl",
		"oth":
		return true
	default:
		return false
	}
}

func verifyPassportID(id string) bool {
	expression := regexp.MustCompile(`^\d{9}$`)
	return len(expression.FindAllString(id, -1)) != 0
}

func validatePassportFields(passport [][]string) bool {
	for _, field := range passport {
		switch field[0] {
		case "byr":
			if !verifyBirthYear(field[1]) {
				return false
			}
		case "iyr":
			if !verifyIssueYear(field[1]) {
				return false
			}
		case "eyr":
			if !verifyExpYear(field[1]) {
				return false
			}
		case "hgt":
			if !verifyHeight(field[1]) {
				return false
			}
		case "hcl":
			if !verifyHairColor(field[1]) {
				return false
			}
		case "ecl":
			if !verifyEyeColor(field[1]) {
				return false
			}
		case "pid":
			if !verifyPassportID(field[1]) {
				return false
			}
		case "cid":
			break
		default:
			return false
		}
	}
	return true
}

func part2(lines []string) int {
	passports := findValidPassports(lines)
	validCount := 0
	for _, passport := range passports {
		if validatePassportFields(passport) {
			validCount++
		}
	}
	return validCount
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
