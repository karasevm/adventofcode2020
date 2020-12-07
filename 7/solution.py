import sys
import re


class Bag:
    def __init__(self):
        self.parents = []
        self.children = {}


def find_unique_parents(bags, target_bag):
    unique_parents = set()
    for parent in bags[target_bag].parents:
        unique_parents.add(parent)
        unique_parents = unique_parents.union(
            find_unique_parents(bags, parent))
    return unique_parents


def find_total_children_count(bags, target_bag):
    total_count = 0
    for child in bags[target_bag].children:
        total_count += bags[target_bag].children[child] * \
            (find_total_children_count(bags, child) + 1)
    return total_count


def parse_bags(lines):
    bags = {}
    expression = re.compile(r"(\d \w+\s\w+)\sbags?")

    for line in lines:
        split_line = line.split(" bags contain ")

        children = expression.findall(split_line[1])
        if not split_line[0] in bags:
            bags[split_line[0]] = Bag()

        for child in children:
            bags[split_line[0]].children[child[2:]] = int(child[:1])
            if not child[2:] in bags:
                bags[child[2:]] = Bag()
            bags[child[2:]].parents.append(split_line[0])
    return bags


def part1(lines):
    return len(find_unique_parents(parse_bags(lines), "shiny gold"))


def part2(lines):
    return find_total_children_count(parse_bags(lines), "shiny gold")


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
