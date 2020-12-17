package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func padState(state [][]string) [][]string {
	// pad in x dimension
	for _, layer := range state {
		for index, line := range layer {
			layer[index] = "." + line + "."
		}
	}
	// pad in z dimension
	for index := range state {
		state[index] = append([]string{strings.Repeat(".", len(state[index][0]))}, state[index]...)
		state[index] = append(state[index], strings.Repeat(".", len(state[index][0])))
	}
	// pad in y dimension
	var frontLayer []string
	var backLayer []string
	for range state[0] {
		temp := strings.Repeat(".", len(state[0][0]))
		frontLayer = append(frontLayer, temp)
		backLayer = append(backLayer, temp)
	}
	state = append([][]string{frontLayer}, state...)
	state = append(state, backLayer)
	return state
}

func doCycle(state [][]string) [][]string {
	var result [][]string
	for i, layer := range state {
		var outputLayer []string
		for j, line := range layer {
			var outputLine string
			for k, symbol := range line {
				count := 0
				for offsetI := -1; offsetI < 2; offsetI++ {
					for offsetJ := -1; offsetJ < 2; offsetJ++ {
						for offsetK := -1; offsetK < 2; offsetK++ {
							if i+offsetI < 0 ||
								j+offsetJ < 0 ||
								k+offsetK < 0 ||
								i+offsetI >= len(state) ||
								j+offsetJ >= len(state[0]) ||
								k+offsetK >= len(state[0][0]) ||
								string(state[i+offsetI][j+offsetJ][k+offsetK]) == "." {
								continue
							}
							count++
						}
					}
				}
				if ((string(symbol) == "#") && (count == 3 || count == 4)) ||
					((string(symbol) == ".") && (count == 3)) {
					outputLine += "#"
				} else {
					outputLine += "."
				}
			}
			outputLayer = append(outputLayer, outputLine)
		}
		result = append(result, outputLayer)
	}
	return result
}

func countActive(state [][]string) int {
	sum := 0
	for _, layer := range state {
		for _, line := range layer {
			for _, symbol := range line {
				if string(symbol) == "#" {
					sum++
				}
			}
		}
	}
	return sum
}

func printState(state [][]string) {
	for i := 0; i < len(state); i++ {
		fmt.Println("layer", i)
		for j := 0; j < len(state[0]); j++ {
			fmt.Println(state[i][j])
		}
	}
}

func part1(lines []string) int {
	state := [][]string{lines}
	for i := 0; i < 6; i++ {
		state = padState(state)
		state = doCycle(state)
	}
	return countActive(state)
}

func padState4d(state [][][]string) [][][]string {
	// pad each cube
	for indexW, cube := range state {
		state[indexW] = padState(cube)
	}
	// pad in w dimension
	var frontCube [][]string
	for range state[0] {
		var layer []string
		for range state[0][0] {
			temp := strings.Repeat(".", len(state[0][0][0]))
			layer = append(layer, temp)
		}
		frontCube = append(frontCube, layer)
	}
	var backCube [][]string
	for range state[0] {
		var layer []string
		for range state[0][0] {
			temp := strings.Repeat(".", len(state[0][0][0]))
			layer = append(layer, temp)
		}
		backCube = append(backCube, layer)
	}
	state = append([][][]string{frontCube}, state...)
	state = append(state, backCube)
	return state
}

func doCycle4d(state [][][]string) [][][]string {
	var result [][][]string
	for w, cube := range state {
		var outputCube [][]string
		for i, layer := range cube {
			var outputLayer []string
			for j, line := range layer {
				var outputLine string
				for k, symbol := range line {
					count := 0
					for offsetW := -1; offsetW < 2; offsetW++ {
						for offsetI := -1; offsetI < 2; offsetI++ {
							for offsetJ := -1; offsetJ < 2; offsetJ++ {
								for offsetK := -1; offsetK < 2; offsetK++ {
									if i+offsetI < 0 ||
										j+offsetJ < 0 ||
										k+offsetK < 0 ||
										w+offsetW < 0 ||
										w+offsetW >= len(state) ||
										i+offsetI >= len(state[0]) ||
										j+offsetJ >= len(state[0][0]) ||
										k+offsetK >= len(state[0][0][0]) ||
										string(state[w+offsetW][i+offsetI][j+offsetJ][k+offsetK]) == "." {
										continue
									}
									count++
								}
							}
						}
					}
					if ((string(symbol) == "#") && (count == 3 || count == 4)) ||
						((string(symbol) == ".") && (count == 3)) {
						outputLine += "#"
					} else {
						outputLine += "."
					}
				}
				outputLayer = append(outputLayer, outputLine)
			}
			outputCube = append(outputCube, outputLayer)
		}
		result = append(result, outputCube)
	}
	return result
}

func countActive4d(state [][][]string) int {
	sum := 0
	for _, w := range state {
		for _, layer := range w {
			for _, line := range layer {
				for _, symbol := range line {
					if string(symbol) == "#" {
						sum++
					}
				}
			}
		}
	}
	return sum
}

func printState4d(state [][][]string) {
	for _, w := range state {
		printState(w)
	}
}

func part2(lines []string) int {
	state := [][][]string{[][]string{lines}}
	for i := 0; i < 6; i++ {
		state = padState4d(state)
		state = doCycle4d(state)
	}
	return countActive4d(state)
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
