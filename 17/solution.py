import sys
import copy


def pad_state(state):
    for layer in state:
        for index, line in enumerate(layer):
            layer[index] = '.'+line+'.'
    for index, _ in enumerate(state):
        state[index].insert(0, '.'*len(state[0][0]))
        state[index].append('.'*len(state[0][0]))
    layer = []
    for _ in state[0]:
        layer.append('.'*len(state[0][0]))
    state.insert(0, copy.deepcopy(layer))
    state.append(copy.deepcopy(layer))
    return state


def do_cycle(state):
    result = []
    for i, layer in enumerate(state):
        output_layer = []
        for j, line in enumerate(layer):
            output_line = ""
            for k, symbol in enumerate(line):
                count = 0
                for offset_i in range(-1, 2):
                    for offset_j in range(-1, 2):
                        for offset_k in range(-1, 2):
                            if ((i + offset_i < 0) or
                                    (j + offset_j < 0) or
                                    (k + offset_k < 0) or
                                    (i + offset_i >= len(state)) or
                                    (j + offset_j >= len(state[0])) or
                                    (k + offset_k >= len(state[0][0])) or
                                    (state[i+offset_i][j+offset_j][k+offset_k] == '.')):
                                continue
                            count += 1
                if (((symbol == '#') and (count in (3, 4))) or
                        ((symbol == '.') and (count == 3))):
                    output_line += '#'
                else:
                    output_line += '.'

            output_layer.append(output_line)
        result.append(output_layer)
    return result


def count_active(state):
    result = 0
    for layer in state:
        for line in layer:
            for symbol in line:
                if symbol == '#':
                    result += 1
    return result


def part1(lines):
    state = [lines]
    for _ in range(6):
        state = pad_state(state)
        state = do_cycle(state)
    return count_active(state)


def pad_state_4d(state):
    for index, cube in enumerate(state):
        state[index] = pad_state(cube)

    empty_cube = []
    layer = []
    for _ in state[0][0]:
        layer.append('.'*len(state[0][0][0]))
    for _ in state[0]:
        empty_cube.append(copy.deepcopy(layer))
    state.insert(0, copy.deepcopy(empty_cube))
    state.append(copy.deepcopy(empty_cube))
    return state


def do_cycle_4d(state):
    result = []
    for w, cube in enumerate(state):
        output_cube = []
        for i, layer in enumerate(cube):
            output_layer = []
            for j, line in enumerate(layer):
                output_line = ""
                for k, symbol in enumerate(line):
                    count = 0
                    for offset_w in range(-1, 2):
                        for offset_i in range(-1, 2):
                            for offset_j in range(-1, 2):
                                for offset_k in range(-1, 2):
                                    if ((i + offset_i < 0) or
                                            (j + offset_j < 0) or
                                            (k + offset_k < 0) or
                                            (w + offset_w < 0) or
                                            (i + offset_i >= len(state[0])) or
                                            (j + offset_j >= len(state[0][0])) or
                                            (k + offset_k >= len(state[0][0][0])) or
                                            (w + offset_w >= len(state)) or
                                            (state[w+offset_w][i+offset_i][j+offset_j][k+offset_k] == '.')):
                                        continue
                                    count += 1
                    if (((symbol == '#') and (count in (3, 4))) or
                            ((symbol == '.') and (count == 3))):
                        output_line += '#'
                    else:
                        output_line += '.'

                output_layer.append(output_line)
            output_cube.append(output_layer)
        result.append(output_cube)
    return result


def count_active_4d(state):
    result = 0
    for cube in state:
        result += count_active(cube)
    return result


def part2(lines):
    state = [[lines]]
    for _ in range(6):
        state = pad_state_4d(state)
        state = do_cycle_4d(state)
    return count_active_4d(state)


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
