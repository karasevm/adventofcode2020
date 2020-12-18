import sys


def solve_expression_ltr(expression: str):
    result = 0
    is_sum = True  # True = addition , False = multiplication
    register = 0
    i = 0
    while i < len(expression):
        if expression[i] in ('+', '*'):
            is_sum = expression[i] == '+'
        elif expression[i] == ' ':
            pass
        else:
            if expression[i] == '(':
                p_count = 1
                j = i + 1
                while p_count != 0:
                    if expression[j] == '(':
                        p_count += 1
                    elif expression[j] == ')':
                        p_count -= 1
                    j += 1
                register = solve_expression_ltr(expression[i+1:j-1])
                i = j
            else:
                register = int(expression[i])
            if is_sum:
                result += register
            else:
                result *= register

        i += 1
    return result


def part1(lines):
    result = 0
    for line in lines:
        result += solve_expression_ltr(line)
    return result


def solve_expression_inverse(expression: str):
    result = 0
    register = 0
    multipliers = []
    i = 0
    while i < len(expression):
        if expression[i] in ('+', '*'):
            if expression[i] == '*':
                multipliers.append(solve_expression_inverse(expression[i+2:]))
                break
        elif expression[i] != ' ':
            if expression[i] == '(':
                p_count = 1
                j = i + 1
                while p_count != 0:
                    if expression[j] == '(':
                        p_count += 1
                    elif expression[j] == ')':
                        p_count -= 1
                    j += 1
                register = solve_expression_inverse(expression[i+1:j-1])
                i = j
            else:
                register = int(expression[i])
            result += register

        i += 1
    for multiplier in multipliers:
        result *= multiplier
    return result


def part2(lines):

    result = 0
    for line in lines:
        result += solve_expression_inverse(line)
    return result


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
