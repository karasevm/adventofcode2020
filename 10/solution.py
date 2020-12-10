import sys


def part1(numbers):
    diff_1_count = 0
    diff_3_count = 0

    sorted_numbers = sorted(numbers)
    sorted_numbers = [0] + sorted_numbers + [sorted_numbers[-1]+3]

    for i in range(len(sorted_numbers)-1):
        if sorted_numbers[i+1] - sorted_numbers[i] == 1:
            diff_1_count += 1
        elif sorted_numbers[i+1] - sorted_numbers[i] == 3:
            diff_3_count += 1

    return diff_1_count * diff_3_count


def count_ways(number):
    if number == 0:
        return 0
    elif number == 1:
        return 1
    elif number == 2:
        return 2
    elif number == 3:
        return 4
    else:
        return count_ways(number-1) + count_ways(number-2) + count_ways(number-3)


def part2(numbers):
    sorted_numbers = sorted(numbers)
    sorted_numbers = [0] + sorted_numbers + [sorted_numbers[-1]+3]

    result = 1
    one_count = 0

    for i in range(len(sorted_numbers)-1):
        if sorted_numbers[i+1] - sorted_numbers[i] == 1:
            one_count += 1
        elif sorted_numbers[i+1] - sorted_numbers[i] == 3:
            if one_count > 0:
                result *= count_ways(one_count)
            one_count = 0

    return result


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
        f"Part 1 answer:{part1(input_numbers)} Part 2 answer: {part2(input_numbers)}")
