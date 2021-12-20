import sys


def part1(lines):
    allergen_map = {}
    all_ingredients = []
    for line in lines:
        ingredients, allergens = line[:-1].split(' (contains ')
        ingredients_list = ingredients.split(' ')
        all_ingredients.extend(ingredients_list)
        allergens_list = allergens.split(', ')
        for allergen in allergens_list:
            if allergen in allergen_map.keys():
                allergen_map[allergen] &= set(ingredients_list)
            else:
                allergen_map[allergen] = set(ingredients_list)
    for k, v in allergen_map.items():
        all_ingredients = [
            ingredient for ingredient in all_ingredients if ingredient not in v]
    return len(all_ingredients)


def part2(lines):
    allergen_map: dict[str, set[str]] = {}
    for line in lines:
        ingredients, allergens = line[:-1].split(' (contains ')
        ingredients_list = ingredients.split(' ')
        allergens_list = allergens.split(', ')
        for allergen in allergens_list:
            if allergen in allergen_map.keys():
                allergen_map[allergen] &= set(ingredients_list)
            else:
                allergen_map[allergen] = set(ingredients_list)
    canonical_map = {}
    while len(canonical_map) != len(allergen_map):
        for allergen, ingredients in allergen_map.items():
            if len(ingredients) == 1:
                canonical_map[allergen] = ingredients.pop()
        for allergen, ingredients in allergen_map.items():
            for _, ingredient in canonical_map.items():
                ingredients -= set([ingredient])
    result = ""
    for key in sorted(canonical_map.keys()):
        result += canonical_map[key] + ","
    return result[:-1]


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
