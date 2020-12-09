import sys


def find_pair(preamble, pair_sum):
    for p1 in preamble:
        for p2 in preamble:
            if p2 != p1 and pair_sum-p1 == p2:
                return True
    return False


def part1(numbers, preamble_size):
    preamble = []
    for index, item in enumerate(numbers[preamble_size:]):
        preamble = numbers[index:index + preamble_size]
        if not find_pair(preamble, item):
            return item
    return 0


def part2(numbers, preamble_size):
    bad_number = part1(numbers, preamble_size)
    for i in range(len(numbers)):
        for k in range(i, len(numbers)-i):
            temp_sum = sum(numbers[i:k])
            if temp_sum > bad_number:
                break
            if temp_sum == bad_number:
                new_numbers = sorted(numbers[i:k])
                return new_numbers[0] + new_numbers[-1]
    return -1


try:
    f = open(sys.argv[1], "r")
except IOError:
    print("Error opening the file, try again")
    sys.exit(1)
with f:
    lines = f.readlines()
    f.close()
    input_numbers = [int(s) for s in lines]
    print(
        f"Part 1 answer:{part1(input_numbers, 25)} Part 2 answer: {part2(input_numbers, 25)}")
