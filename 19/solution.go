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

func parseRules(rules []string) map[int]string {
	result := make(map[int]string)
	for _, rule := range rules {
		splitIndex := strings.Index(rule, ":")
		tempKey, err := strconv.Atoi(rule[:splitIndex])
		if err != nil {
			log.Fatal(err)
		}
		result[tempKey] = rule[splitIndex+2:]
	}
	return result
}

// generates all strings that match the provided set of rules
func checkRule(rules map[int]string, rule string) []string {
	// list of lists of results, length depends on rule count (1 or 2)
	var resultParts [][]string
	// for each rule part (separated by spaces)
	ruleParts := strings.Split(rule, " ")
	for _, rulePart := range ruleParts {
		// if rule is a letter, return that letter
		if string(rulePart[0]) == "\"" {
			return []string{string(rulePart[1])}
		}

		// parse rule part to int
		rulePartID, err := strconv.Atoi(rulePart)
		if err != nil {
			log.Fatal(err)
		}

		receivedStrings := make([]string, 0)

		// get rule string referred by the current rule, split them if they are OR'd
		rulesToCheck := strings.Split(rules[rulePartID], " | ")
		// check those rules recursively, there can be maximum 2 rules
		for _, ruleToCheck := range rulesToCheck {
			receivedStrings = append(receivedStrings, checkRule(rules, ruleToCheck)...)
		}
		// push result of current rule to resultParts
		resultParts = append(resultParts, receivedStrings)
	}

	var result []string

	// if current rule consists of 2 parts, create all possible combinations
	if len(resultParts) == 2 {
		for _, firstPart := range resultParts[0] {
			for _, secondPart := range resultParts[1] {
				result = append(result, firstPart+secondPart)
			}
		}
	} else {
		result = resultParts[0]
	}

	return result
}

// generates a regex from a provided list of rules
func checkRecursiveRule(rules map[int]string, rule string) string {
	// list of lists of results, length depends on rule count
	var resultParts [][]string
	// for each rule part (separated by spaces)
	ruleParts := strings.Split(rule, " ")
	for _, rulePart := range ruleParts {
		// if rule is a letter, return that letter
		if string(rulePart[0]) == "\"" {
			return string(rulePart[1])
		}

		// parse rule part to int
		rulePartID, err := strconv.Atoi(rulePart)
		if err != nil {
			log.Fatal(err)
		}

		receivedStrings := make([]string, 0)

		// get rule string referred by the current rule, split them if they are OR'd
		rulesToCheck := strings.Split(rules[rulePartID], " | ")
		// check those rules recursively, there can be maximum 2 rules
		for _, ruleToCheck := range rulesToCheck {
			if ruleToCheck == "42 8" {
				specialString42 := checkRecursiveRule(rules, "42")
				receivedStrings = append(receivedStrings, "("+specialString42+")+")
			} else if ruleToCheck == "42 11 31" {
				specialString42 := checkRecursiveRule(rules, "42")
				specialString31 := checkRecursiveRule(rules, "31")
				receivedStrings = append(receivedStrings, "("+specialString42+"){XXXX}"+"("+specialString31+"){XXXX}")
			} else if ruleToCheck != "42 31" {
				receivedStrings = append(receivedStrings, checkRecursiveRule(rules, ruleToCheck))
			}
		}
		// push result of current rule to resultParts
		resultParts = append(resultParts, receivedStrings)
	}

	var result = "("
	for part := range resultParts {
		result += "("
		for _, firstPart := range resultParts[part] {
			result += firstPart
			result += "|"
		}
		result = strings.TrimSuffix(result, "|")
		result += ")"
	}
	result = strings.TrimSuffix(result, "|")
	result += ")"

	return result

}

func part1(lines []string) int {
	// find the rule/input separator
	ruleEndIndex := 0
	for index, line := range lines {
		if line == "" {
			ruleEndIndex = index
			break
		}
	}

	// get all valid strings
	var rules = parseRules(lines[:ruleEndIndex])
	validStrings := checkRule(rules, "0")

	// check how many strings are in valid string slice (o(n^2) but not that important)
	count := 0
	for _, line := range lines[ruleEndIndex+1:] {
		for _, validString := range validStrings {
			if line == validString {
				count++
				break
			}
		}
	}

	return count
}

func part2(lines []string) int {
	// find the rule/input separator
	ruleEndIndex := 0
	for index, line := range lines {
		if line == "" {
			ruleEndIndex = index
			break
		}
	}

	// get all valid strings
	var rules = parseRules(lines[:ruleEndIndex])
	// adjust input
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"

	validStringRegex := "^" + checkRecursiveRule(rules, "0") + "$"

	count := 0
	for _, line := range lines[ruleEndIndex+1:] {
		for i := 1; i < 10; i++ {
			var currentRegex = strings.ReplaceAll(validStringRegex, "XXXX", fmt.Sprintf("%v", i))
			matched, _ := regexp.MatchString(currentRegex, line)
			if matched {
				count++
				break
			}
		}
	}

	return count
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
