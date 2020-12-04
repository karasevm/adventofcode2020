import sys
import re


def parse_passports(lines):
    passports = []
    passport = []
    expression = re.compile(r'(\w{3}):([^\s]+)')
    for line in lines:
        if len(line) == 1:
            passports.append(passport)
            passport = []
            continue
        match = expression.findall(line)
        if len(match) > 0:
            passport = passport + match
    return passports


def find_valid_passports(lines):
    valid_passports = []
    for passport in parse_passports(lines):
        field_count = len(passport)
        # print(passport)
        for field in passport:
            if field[0] == 'cid':
                field_count -= 1
                break
        if field_count >= 7:
            valid_passports.append(passport)
    return valid_passports


def part1(lines):
    return len(find_valid_passports(lines))


def verify_birth_year(year):
    return (int(year) >= 1920) and (int(year) <= 2002)


def verify_issue_year(year):
    return (int(year) >= 2010) and (int(year) <= 2020)


def verify_exp_year(year):
    return (int(year) >= 2020) and (int(year) <= 2030)


def verify_height(height):
    if height[-2:] == 'cm':
        return (int(height[:-2]) >= 150) and (int(height[:-2]) <= 193)
    elif height[-2:] == 'in':
        return (int(height[:-2]) >= 59) and (int(height[:-2]) <= 76)
    else:
        return False


def verify_hair_color(color):
    hex_expression = re.compile(r'\#[0-9a-z]{6}')
    return len(hex_expression.findall(color)) != 0


def verify_eye_color(color):
    colors = {'amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'}
    return color in colors


def verify_passport_ID(passport_id):
    expression = re.compile(r'^\d{9}$')
    return len(expression.findall(passport_id)) != 0


def valid_passport_fields(passport):
    for field in passport:
        if field[0] == 'byr':
            if not verify_birth_year(field[1]):
                return False
        elif field[0] == 'iyr':
            if not verify_issue_year(field[1]):
                return False
        elif field[0] == 'eyr':
            if not verify_exp_year(field[1]):
                return False
        elif field[0] == 'hgt':
            if not verify_height(field[1]):
                return False
        elif field[0] == 'hcl':
            if not verify_hair_color(field[1]):
                return False
        elif field[0] == 'ecl':
            if not verify_eye_color(field[1]):
                return False
        elif field[0] == 'pid':
            if not verify_passport_ID(field[1]):
                return False
        elif field[0] == 'cid':
            pass
        else:
            return False
    return True


def part2(lines):
    passports = find_valid_passports(lines)
    valid_count = 0
    for passport in passports:
        if valid_passport_fields(passport):
            valid_count += 1
    return valid_count


try:
    f = open(sys.argv[1], 'r')
except IOError:
    print("Error opening the file, try again")
    sys.exit(1)
with f:
    rawLines = f.readlines()
    f.close()
    print(
        f"Part 1 answer:{part1(rawLines)} Part 2 answer: {part2(rawLines)}")
