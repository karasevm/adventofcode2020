import sys
from collections import defaultdict, deque
import math


def parse_tiles(lines: list[str]) -> dict[int, list[str]]:
    result_tiles = {}
    temp_grid = []
    prev_title = ''
    for line_number, line in enumerate(lines):
        # if line[:4] == "Tile" and len(temp_grid) != 0:
        if len(line) == 0:
            result_tiles[int(prev_title)] = temp_grid
            temp_grid = []
        elif line[:4] == "Tile":
            prev_title = line[5:-1]
        elif len(line) != 0:  # read single tile
            temp_grid.append(line)
    return result_tiles


def check_sides_overlap(grid_a: list[str], grid_b: list[str]) -> bool:

    offsets = [0, 1, 2, 3, 10, 11, 12, 13]
    for offset in offsets:
        rotated_grid_b = rotate_grid_clockwise(grid_b, offset)
        if grid_a[0] == rotated_grid_b[-1]:
            return True
        if grid_a[-1] == rotated_grid_b[0]:
            return True
        if ''.join([row[0] for row in grid_a]) == ''.join([row[-1] for row in rotated_grid_b]):
            return True
        if ''.join([row[-1] for row in grid_a]) == ''.join([row[0] for row in rotated_grid_b]):
            return True
    return False


def part1(lines):
    # tile_borders - dict of tiles, containing lists of possible rotations,
    # containing lists of resulting sides
    tiles = parse_tiles(lines)
    result = 1
    for tile_id, tile in tiles.items():
        count = 0
        for other_tile_id, other_tile in tiles.items():
            overlap = check_sides_overlap(tile, other_tile)
            if other_tile_id != tile_id and overlap:
                count += 1
        if count == 2:  # if tile has 2 neighbors
            result *= tile_id
    return result


def rotate_grid_clockwise(grid: list[str], rotation: int) -> list[str]:
    if rotation > 3:
        grid = [line[::-1] for line in grid]
        rotation = rotation - 10
    for _ in range(rotation):
        grid = [''.join(x) for x in zip(*grid[::-1])]
    return grid


def trim_grid(grid: list[str]) -> list[str]:
    grid = grid[1:-1]
    grid = [row[1:-1] for row in grid]
    return grid


def print_grid(grid: list[str]) -> None:
    for line in grid:
        print(line)


def is_monster(grid: list[str], i: int, j: int) -> bool:
    monster_offsets = [(0, 18), (1, 0), (1, 5), (1, 6), (1, 11), (1, 12),
                       (1, 17), (1, 18), (1, 19), (2, 1), (2, 4), (2, 7), (2, 10), (2, 16)]
    for offset_i, offset_j in monster_offsets:
        if grid[i+offset_i][j+offset_j] != "#":
            return False
    print(f"Found monster at {i} {j}")
    return True


def count_monsters(grid: list[str]) -> int:
    count = 0
    for i in range(1, len(grid)-1):
        for j in range(len(grid[0])-19):
            if is_monster(grid, i, j):
                count += 1
    return count


def part2(lines):
    tile_grids = parse_tiles(lines)

    tile_neighbors = defaultdict(list[int])
    for tile_id, tile in tile_grids.items():
        for other_tile_id, other_tile in tile_grids.items():
            overlap = check_sides_overlap(tile, other_tile)
            if other_tile_id != tile_id and overlap:
                tile_neighbors[tile_id].append(other_tile_id)

    grid_size = int(math.sqrt(len(tile_grids)))*3
    result_grid: list[list[list[str]]] = [[[]
                                           for _ in range(grid_size)] for _ in range(grid_size)]
    grid_center = grid_size // 2

    # put first occurring tile with 4 neighbors in the center
    for tile, neighbors in tile_neighbors.items():
        if len(neighbors) == 4:
            result_grid[grid_center][grid_center] = tile_grids.pop(tile)
            break

    grid_queue = deque()
    grid_queue.append((grid_center, grid_center))
    cycle_count = 0
    visited = set()
    while grid_queue:
        cycle_count += 1
        row, col = grid_queue.popleft()
        current_grid = result_grid[row][col]
        for tile_id, tile_grid in tile_grids.items():

            if tile_id in visited:
                continue

            offsets = [0, 1, 2, 3, 10, 11, 12, 13]
            for offset in offsets:
                rotated_tile_grid = rotate_grid_clockwise(tile_grid, offset)
                if current_grid[0] == rotated_tile_grid[-1]:  # top
                    result_grid[row-1][col] = rotated_tile_grid
                    grid_queue.append((row-1, col))
                    visited.add(tile_id)
                elif current_grid[-1] == rotated_tile_grid[0]:  # bottom
                    result_grid[row+1][col] = rotated_tile_grid
                    grid_queue.append((row+1, col))
                    visited.add(tile_id)
                elif ''.join([row[0] for row in current_grid]) == ''.join([row[-1] for row in rotated_tile_grid]):  # left
                    result_grid[row][col-1] = rotated_tile_grid
                    grid_queue.append((row, col-1))
                    visited.add(tile_id)
                elif ''.join([row[-1] for row in current_grid]) == ''.join([row[0] for row in rotated_tile_grid]):  # right
                    result_grid[row][col+1] = rotated_tile_grid
                    grid_queue.append((row, col+1))
                    visited.add(tile_id)

    result_start_row = sys.maxsize
    result_start_col = sys.maxsize
    for row_index, row in enumerate(result_grid):
        for col_index, col in enumerate(row):
            if len(col) != 0:
                if row_index < result_start_row:
                    result_start_row = row_index
                if col_index < result_start_col:
                    result_start_col = col_index

    combined_grid = []
    for row in result_grid[result_start_row:result_start_row+grid_size//3]:
        for i in range(8):
            tmp_line = ""
            for col in row[result_start_col: result_start_col+grid_size//3]:
                if len(col) != 0:
                    tmp_line += trim_grid(col)[i]
            combined_grid.append(tmp_line)
    print_grid(combined_grid)

    monster_count = 0
    while monster_count == 0:
        offsets = [0, 1, 2, 3, 10, 11, 12, 13]
        for offset in offsets:
            monster_count += count_monsters(
                rotate_grid_clockwise(combined_grid, offset))

    result_hash_count = 0
    for line in combined_grid:
        result_hash_count += line.count("#")

    result_hash_count -= monster_count * 15
    return result_hash_count


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
