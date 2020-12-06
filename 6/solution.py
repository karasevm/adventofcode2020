import sys


def count_unique(lines):
    letters = set()
    count = 0
    for line in lines:
        if line == "\n":
            count += len(letters)
            letters = set()
            continue

        for letter in line:
            if letter != '\n':
                letters.add(letter)
    count += len(letters)
    return count


def part1(lines):
    return count_unique(lines)


def count_common(lines):
    line_group = []
    letters = set()
    count = 0
    for line in lines:
        if line == "\n":
            count += len(line_group[0].intersection(*line_group[1:]))
            line_group = []
            continue
        letters = set()
        for letter in line:
            if letter != '\n':
                letters.add(letter)
        line_group.append(letters)
    count += len(line_group[0].intersection(*line_group[1:]))
    return count


def part2(lines):
    return count_common(lines)


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
