import sys
from collections import defaultdict


def get_tile_map(lines: list[str]) -> defaultdict[tuple[int, int], bool]:
    tile_map = defaultdict(bool)
    for line in lines:
        i = 0
        x = 0
        y = 0
        while i < len(line):
            match line[i]:
                case "e":
                    x += 1
                case "w":
                    x -= 1
                case "n":
                    i += 1
                    y -= 1
                    match line[i]:
                        case "e":
                            x += 1
                        case "w":
                            pass
                case "s":
                    i += 1
                    y += 1
                    match line[i]:
                        case "e":
                            pass
                        case "w":
                            x -= 1
            i += 1
        tile_map[(x, y)] = not tile_map[(x, y)]

    return tile_map


def part1(lines: list[str]):
    return list(get_tile_map(lines).values()).count(True)


def part2(lines):
    tiles = get_tile_map(lines)
    for _ in range(100):
        new_tiles = defaultdict(bool)
        for x in range(-100, 100):
            for y in range(-100, 100):
                black_tile_count = 0
                offsets = [(-1, 0), (0, -1), (1, -1), (1, 0), (0, 1), (-1, 1)]
                for offset in offsets:
                    if tiles[(x+offset[0], y+offset[1])]:
                        black_tile_count += 1
                if tiles[(x, y)] and (black_tile_count == 0 or black_tile_count > 2):
                    new_tiles[(x, y)] = False
                    continue
                if (tiles[(x, y)] == False) and black_tile_count == 2:
                    new_tiles[(x, y)] = True
                    continue
                if tiles[(x, y)]:
                    new_tiles[(x, y)] = tiles[(x, y)]
        tiles = new_tiles
    return list(tiles.values()).count(True)


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
        f"Part 1 answer: {part1(rawLines)} Part 2 answer: {part2(rawLines)}")
