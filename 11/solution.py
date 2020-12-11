import sys
import copy
import math


def count_occupied(x, y, lines):
    directions = {(1, 0), (1, 1), (0, 1), (-1, 1),
                  (-1, 0), (-1, -1), (0, -1), (1, -1)}
    count = 0
    for direction in directions:
        y_to_check = y + direction[0]
        x_to_check = x + direction[1]
        if (y_to_check < 0 or y_to_check >= len(lines) or
                x_to_check < 0 or x_to_check >= len(lines[0])):
            continue
        if lines[y_to_check][x_to_check] == '#':
            count += 1
    return count


def count_total_occupied(lines):
    count = 0
    for line in lines:
        count += line.count('#')
    return count


def part1(lines):
    list_write = copy.deepcopy(lines)
    list_read = []

    while list_read != list_write:
        list_read = list_write
        list_write = []

        for y in range(len(lines)):
            new_line = ""
            for x in range(len(lines[0])):
                if list_read[y][x] == 'L':
                    if count_occupied(x, y, list_read) == 0:
                        new_line = new_line + '#'
                    else:
                        new_line = new_line + 'L'
                elif list_read[y][x] == '#':
                    if count_occupied(x, y, list_read) >= 4:
                        new_line = new_line + 'L'
                    else:
                        new_line = new_line + '#'
                else:
                    new_line = new_line + '.'
            list_write.append(new_line)
    return count_total_occupied(list_write)


def count_occupied_lines(x, y, lines):
    directions = [(1, 0), (1, 1), (0, 1), (-1, 1),
                  (-1, 0), (-1, -1), (0, -1), (1, -1)]
    count = 0
    for multiplier in range(1, math.ceil(math.sqrt(x*x+y*y)) + 1):
        for i in range(len(directions)-1, -1, -1):
            y_to_check = y + directions[i][0] * multiplier
            x_to_check = x + directions[i][1] * multiplier

            if (y_to_check < 0 or y_to_check >= len(lines) or
                    x_to_check < 0 or x_to_check >= len(lines[0])):
                continue

            if (lines[y_to_check][x_to_check] == '#' or
                    lines[y_to_check][x_to_check] == 'L'):
                directions.pop(i)

                if lines[y_to_check][x_to_check] == '#':
                    count += 1
    return count


def part2(lines):
    list_write = copy.deepcopy(lines)
    list_read = []

    while list_read != list_write:
        list_read = list_write
        list_write = []

        for y in range(len(lines)):
            new_line = ""
            for x in range(len(lines[0])):
                if list_read[y][x] == 'L':
                    if count_occupied_lines(x, y, list_read) == 0:
                        new_line = new_line + '#'
                    else:
                        new_line = new_line + 'L'
                elif list_read[y][x] == '#':
                    if count_occupied_lines(x, y, list_read) >= 5:
                        new_line = new_line + 'L'
                    else:
                        new_line = new_line + '#'
                else:
                    new_line = new_line + '.'
            list_write.append(new_line)
    return count_total_occupied(list_write)


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
