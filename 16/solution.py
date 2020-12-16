import sys
import re
import typing


class Field:
    def __init__(self,
                 range_a: typing.Tuple[int, int],
                 range_b: typing.Tuple[int, int],
                 name: str):
        self.min_a = int(range_a[0])
        self.max_a = int(range_a[1])
        self.min_b = int(range_b[0])
        self.max_b = int(range_b[1])
        self.name = name


def parse_fields(lines):
    result = []
    expression = re.compile(r"^((?:\w|\s)+)\:\s(\d+)\-(\d+)\sor\s(\d+)\-(\d+)")
    for line in lines:
        match = expression.findall(line)
        result.append(Field((match[0][1], match[0][2]),
                            (match[0][3], match[0][4]), match[0][0]))
    return result


def value_fits_field(value, field):
    return ((field.min_a <= value <= field.max_a) or
            (field.min_b <= value <= field.max_b))


def value_fits_field_list(value, field_list):
    for field in field_list:
        if value_fits_field(value, field):
            return True
    return False


def part1(lines):
    field_end_index = 0
    nearby_tickets_start = 0
    for index, line in enumerate(lines):
        if line == "your ticket:\n":
            field_end_index = index - 1
            nearby_tickets_start = index + 4
            break

    fields = parse_fields(lines[:field_end_index])
    result_sum = 0
    for line in lines[nearby_tickets_start:]:
        values = line.split(",")
        for val in values:
            if not value_fits_field_list(int(val), fields):
                result_sum += int(val)
    return result_sum


def part2(lines):
    field_end_index = 0
    nearby_tickets_start = 0
    your_ticket_index = 0
    for index, line in enumerate(lines):
        if line == "your ticket:\n":
            field_end_index = index - 1
            your_ticket_index = index + 1
            nearby_tickets_start = index + 4
            break

    fields = parse_fields(lines[:field_end_index])
    clean_tickets = []
    for line in lines[nearby_tickets_start:]:
        values = line.split(",")
        clean_ticket = []
        for val in values:
            if value_fits_field_list(int(val), fields):
                clean_ticket.append(int(val))
        clean_tickets.append(clean_ticket)

    column_possible_fields = [range(len(fields))
                              for j in range(len(clean_tickets[0]))]

    for ticket in clean_tickets:
        for comlumn_index, column in enumerate(ticket):
            fitting_fields = []
            for field_index, field in enumerate(fields):
                if value_fits_field(column, field):
                    fitting_fields.append(field_index)
            column_possible_fields[comlumn_index] = list(
                set(column_possible_fields[comlumn_index]) & set(fitting_fields))

    result = {}
    for column_a in column_possible_fields:
        for index, column_b in enumerate(column_possible_fields):
            if len(column_b) == 1:
                result[index] = fields[column_b[0]].name
            elif len(column_b) == len(column_a) + 1:
                result[index] = fields[list(
                    list(set(column_a)-set(column_b)) + list(set(column_b)-set(column_a)))[0]].name

    result_sum = 1
    your_ticket_split = lines[your_ticket_index].split(",")
    for index, name in result.items():
        if len(name) >= 9 and name[:9] == "departure":
            result_sum *= int(your_ticket_split[index])
    return result_sum


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
