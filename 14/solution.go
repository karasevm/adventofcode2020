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

func setBit(value uint64, pos uint) uint64 {
	value |= (1 << pos)
	return value
}

func clearBit(value uint64, pos uint) uint64 {
	value &= ^(1 << pos)
	return value
}

func applyMaskToVal(inputInt uint64, mask string) uint64 {
	for i := len(mask) - 1; i >= 0; i-- {
		pos := uint(len(mask) - i - 1)
		if string(mask[i]) == "1" {
			inputInt = setBit(inputInt, pos)
		} else if string(mask[i]) == "0" {
			inputInt = clearBit(inputInt, pos)
		}
	}
	return inputInt
}

func part1(lines []string) uint64 {
	var mask string
	memory := make(map[uint64]uint64)
	expression := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		if line[:4] == "mask" {
			mask = line[7:]
		} else {
			match := expression.FindAllString(line, -1)
			addr, err := strconv.ParseUint(match[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.ParseUint(match[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			memory[addr] = applyMaskToVal(value, mask)
		}
	}
	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}
	return sum
}

func applyMaskToAddr(inputInt uint64, mask string) []uint64 {
	result := make(map[uint64]struct{})
	for i := len(mask) - 1; i >= 0; i-- {
		pos := uint(len(mask) - i - 1)
		if string(mask[i]) == "1" {
			inputInt = setBit(inputInt, pos)
		}
	}
	result[inputInt] = struct{}{}
	for i := len(mask) - 1; i >= 0; i-- {
		pos := uint(len(mask) - i - 1)
		if string(mask[i]) == "X" {
			var temp []uint64
			for addr := range result {
				temp = append(temp, addr)
			}
			for _, addr := range temp {
				result[setBit(addr, pos)] = struct{}{}
				result[clearBit(addr, pos)] = struct{}{}
			}
		}
	}
	var resultSlice []uint64
	for k := range result {
		resultSlice = append(resultSlice, k)
	}
	return resultSlice
}

func part2(lines []string) uint64 {
	var mask string
	memory := make(map[uint64]uint64)
	expression := regexp.MustCompile(`\d+`)
	for _, line := range lines {
		if line[:4] == "mask" {
			mask = line[7:]
		} else {
			match := expression.FindAllString(line, -1)
			addr, err := strconv.ParseUint(match[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			value, err := strconv.ParseUint(match[1], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			addresses := applyMaskToAddr(addr, mask)
			for _, maskedAddr := range addresses {
				memory[maskedAddr] = value
			}

		}
	}
	sum := uint64(0)
	for _, v := range memory {
		sum += v
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
