import sys
import re
import math


def set_bit(value, pos):
    return value | (1 << pos)


def clear_bit(value, pos):
    return value & ~(1 << pos)


def apply_mask_to_val(input_int, mask):
    for i in range(len(mask)-1, -1, -1):
        pos = len(mask) - i - 1
        if mask[i] == '1':
            input_int = set_bit(input_int, pos)
        elif mask[i] == '0':
            input_int = clear_bit(input_int, pos)
    return input_int


def part1(lines):
    mask = ''
    memory = {}
    expression = re.compile(r"\d+")
    for line in lines:
        if line[:4] == "mask":
            mask = line[7:-1]
        else:
            match = expression.findall(line)
            memory[int(match[0])] = apply_mask_to_val(int(match[1]), mask)
    return sum(memory.values())


def apply_mask_to_addr(input_int, mask):
    result = set()
    for i in range(len(mask)-1, -1, -1):
        pos = len(mask) - i - 1
        if mask[i] == '1':
            input_int = set_bit(input_int, pos)
    result.add(input_int)
    for i in range(len(mask)-1, -1, -1):
        pos = len(mask) - i - 1
        if mask[i] == 'X':
            temp = list(result)
            for addr in temp:
                result.add(set_bit(addr, pos))
                result.add(clear_bit(addr, pos))
    return result


def part2(lines):
    mask = ''
    memory = {}
    expression = re.compile(r"\d+")
    for line in lines:
        if line[:4] == "mask":
            mask = line[7:-1]
        else:
            match = expression.findall(line)
            addresses = apply_mask_to_addr(int(match[0]), mask)
            for masked_address in addresses:
                memory[masked_address] = int(match[1])

    return sum(memory.values())


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
