import sys


def solve(line: str, until: int):
    current_digit = 0
    digits_last_seen = {}
    start_digit_list = [int(n) for n in line.split(',')]
    for i in range(len(start_digit_list)):
        digits_last_seen[start_digit_list[i]] = i
        current_digit = start_digit_list[i]
    for i in range(len(start_digit_list)-1, until-1):
        age = 0
        if current_digit in digits_last_seen:
            age = i - digits_last_seen[current_digit]
        digits_last_seen[current_digit] = i
        current_digit = age
    return current_digit


def part1(lines):
    return solve(lines[0][:-1], 2020)


def part2(lines):
    return solve(lines[0][:-1], 30000000)


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
