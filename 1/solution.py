import sys


def part1(input_arr):
    for a in input_arr:
        for b in input_arr:
            if a + b == 2020:
                return a * b
    return -1


def part2(input_arr):
    for a in input_arr:
        for b in input_arr:
            for c in input_arr:
                if a + b + c == 2020:
                    return a * b * c
    return -1


try:
    f = open(sys.argv[1], "r")
except IOError:
    print("Error opening the file, try again")
    sys.exit(1)
with f:
    lines = f.readlines()
    f.close()
    digitArray = [int(s) for s in lines]
    print(
        f"Part 1 answer:{part1(digitArray)} Part 2 answer: {part2(digitArray)}")
