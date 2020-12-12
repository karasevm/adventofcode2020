import sys


def part1(lines):
    direction_multipliers = {0: (0, 1), 90: (1, 0), 180: (0, -1), 270: (-1, 0)}
    current_direction = 90
    current_x = 0
    current_y = 0
    for line in lines:
        action = line[:1]
        value = int(line[1:])
        if action == 'N':
            current_y += value
        elif action == 'S':
            current_y -= value
        elif action == 'E':
            current_x += value
        elif action == 'W':
            current_x -= value
        elif action == 'L':
            current_direction = (current_direction - value + 360) % 360
        elif action == 'R':
            current_direction = (current_direction + value) % 360
        elif action == 'F':
            current_x += value * direction_multipliers[current_direction][0]
            current_y += value * direction_multipliers[current_direction][1]
    return abs(current_x) + abs(current_y)


def part2(lines):
    ship_x = 0
    ship_y = 0
    waypoint_x = 10
    waypoint_y = 1
    for line in lines:
        action = line[:1]
        value = int(line[1:])
        if action == 'N':
            waypoint_y += value
        elif action == 'S':
            waypoint_y -= value
        elif action == 'E':
            waypoint_x += value
        elif action == 'W':
            waypoint_x -= value
        elif action == 'R':
            for _ in range(value//90):
                waypoint_x, waypoint_y = waypoint_y, -waypoint_x
        elif action == 'L':
            for _ in range(value//90):
                waypoint_x, waypoint_y = -waypoint_y, waypoint_x
        elif action == 'F':
            ship_x += waypoint_x * value
            ship_y += waypoint_y * value
    return abs(ship_x) + abs(ship_y)


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
