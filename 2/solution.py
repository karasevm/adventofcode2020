import sys
import re


def part1(lines):
    valid_count = 0
    expression = re.compile(r'(\d+)\-(\d+)\ ([a-z]):\ ([a-z]+)$')
    for line in lines:
        match = expression.findall(line)
        count = match[0][3].count(match[0][2])
        if count >= int(match[0][0]) and count <= int(match[0][1]):
            valid_count += 1
    return valid_count


def part2(lines):
    valid_count = 0
    expression = re.compile(r'(\d+)\-(\d+)\ ([a-z]):\ ([a-z]+)$')
    for line in lines:
        match = expression.findall(line)
        if ((match[0][2] == match[0][3][int(match[0][0])-1] and match[0][2] != match[0][3][int(match[0][1])-1]) or
                (match[0][2] != match[0][3][int(match[0][0])-1] and match[0][2] == match[0][3][int(match[0][1])-1])):
            valid_count += 1
    return valid_count


try:
    f = open(sys.argv[1], "r")
except IOError:
    print("Error opening the file, try again")
    sys.exit(1)
with f:
    rawLines = f.readlines()
    f.close()
    #digitArray = [int(s) for s in lines]
    print(
        f"Part 1 answer:{part1(rawLines)} Part 2 answer: {part2(rawLines)}")
