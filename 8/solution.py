import sys


def execute_operations(lines):
    hit_instructions = set()
    accumulator = 0
    i = 0
    while i < len(lines):
        if i in hit_instructions:
            return 1, accumulator
        hit_instructions.add(i)
        if lines[i][:3] == "jmp":
            i += int(lines[i][4:])
            continue
        elif lines[i][:3] == "acc":
            accumulator += int(lines[i][4:])
        i += 1
    return 0, accumulator


def part1(lines):
    return execute_operations(lines)[1]


def part2(lines):
    for i in range(len(lines)):
        temp_line = ""
        if lines[i][:3] == "jmp":
            temp_line = lines[i].replace("jmp", "nop")
        elif lines[i][:3] == "acc":
            temp_line = lines[i].replace("nop", "jmp")
        status, accumulator = execute_operations(
            lines[:i] + [temp_line] + lines[i+1:])
        if status == 0:
            return accumulator
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
