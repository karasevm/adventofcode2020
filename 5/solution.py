import sys
import math


def parse_bsp(bspString):
    row_start = 0
    row_end = 127
    col_start = 0
    col_end = 7

    for character in bspString[:7]:
        pivot = math.ceil((row_start + row_end)/2)
        if character == 'B':
            row_start = pivot
        else:
            row_end = pivot - 1

    for character in bspString[7:]:
        pivot = math.ceil((col_start + col_end) / 2)
        if character == 'R':
            col_start = pivot
        else:
            col_end = pivot
    return (row_start,
            col_start,
            (row_start * 8) + col_start)


def part1(lines):
    max_id = -1
    for line in lines:
        result = parse_bsp(line)
        if max_id < result[2]:
            max_id = result[2]
    return max_id


def part2(lines):
    seats = []
    for line in lines:
        seats.append(parse_bsp(line))
    seats.sort(key=lambda tup: tup[2])
    for i in range(0, len(seats)-1):
        if seats[i][2]+2 == seats[i+1][2]:
            return seats[i][2]+1

    return -1


try:
    f = open(sys.argv[1], "r")
except IOError:
    print("Error opening the file, try again")
    sys.exit(1)
with f:
    rawLines = f.readlines()
    f.close()
    # digitArray = [int(s) for s in lines]
    print(
        f"Part 1 answer:{part1(rawLines)} Part 2 answer: {part2(rawLines)}")
