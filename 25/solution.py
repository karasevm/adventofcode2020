import sys


def transform_key(public_key, loop_size):
    subject_number = 1
    for _ in range(loop_size):
        subject_number = (subject_number * public_key) % 20201227
    return subject_number


def find_loop_size(public_key):
    subject_number = 1
    loop_size = 0
    while subject_number != public_key:
        subject_number = (subject_number * 7) % 20201227
        loop_size += 1
    return loop_size


def part1(lines):
    # print(find_loop_size(5764801))
    return transform_key(int(lines[0]), find_loop_size(int(lines[1])))


def part2(lines):
    return 0


try:
    f = open(sys.argv[1], "r")
except IOError:
    print("Error opening the file, try again")
    sys.exit(1)
with f:
    rawLines = f.read().splitlines()
    f.close()
    # digitArray = [int(s) for s in lines]
    print(
        f"Part 1 answer:{part1(rawLines)} Part 2 answer: {part2(rawLines)}")
