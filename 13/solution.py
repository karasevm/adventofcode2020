import sys
import re
import math


def part1(lines):
    expression = re.compile(r"\d{1,4}")
    intervals = [int(s) for s in expression.findall(lines[1])]
    start_time = int(lines[0])
    min_diff = sys.maxsize
    min_diff_interval = 0
    for interval in intervals:
        multiplier = 1
        while interval * multiplier < start_time:
            multiplier += 1
        current_diff = (interval * multiplier) - start_time
        if current_diff < min_diff:
            min_diff = current_diff
            min_diff_interval = interval
    return min_diff * min_diff_interval


def lcm(a, b):
    return abs(a*b) // math.gcd(a, b)


def part2(lines):
    expression = re.compile(r"\d{1,4}|x")
    intervals = [1 if s == 'x' else int(s)
                 for s in expression.findall(lines[1])]
    next_multipliers = []
    multiplier = 0
    multiplier_offset = 1
    multiplier_step = 1
    for i in range(1, len(intervals)):
        next_multipliers = []
        multiplier = multiplier_offset
        while len(next_multipliers) < 2:
            multiple = int(intervals[i-1] * multiplier)
            if (multiple+1) % intervals[i] == 0:
                next_multipliers.append(int((multiple+1)/intervals[i]))
            multiplier += multiplier_step
        multiplier_step = lcm(
            multiplier_step, next_multipliers[1] - next_multipliers[0])
        multiplier_offset = next_multipliers[0]
    return next_multipliers[0] * intervals[-1] - len(intervals) + 1


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
