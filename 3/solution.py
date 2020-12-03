import sys


def tree_finder(grid, right, down):
    tree_count = 0
    x = 0
    y = 0
    line_length = len(grid[0]) - 1
    while y < len(grid):
        if grid[y][x] == "#":
            tree_count += 1
        x = (x + right) % line_length
        y += down
    return tree_count


def part1(lines):
    return tree_finder(lines, 3, 1)


def part2(lines):
    return (tree_finder(lines, 1, 1) *
            tree_finder(lines, 3, 1) *
            tree_finder(lines, 5, 1) *
            tree_finder(lines, 7, 1) *
            tree_finder(lines, 1, 2))


try:
    f = open(sys.argv[1], "r")
except IOError:
    print("Error opening the file, try again")
    sys.exit(1)
with f:
    rawLines = f.readlines()
    f.close()
    print(
        f"Part 1 answer:{part1(rawLines)} Part 2 answer: {part2(rawLines)}")
