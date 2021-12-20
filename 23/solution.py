import sys


def part1(lines):
    cups = [int(n) for n in lines[0]]
    current_cup = cups[0]
    current_cup_index = 0
    temp_cups = list(cups)
    for _ in range(100):
        print("============================")
        print(temp_cups)
        print("current cup", current_cup)
        destination_cup = current_cup-1
        removed_cups = []
        for i in range(3):
            temp_index = (current_cup_index+1)
            if temp_index >= len(temp_cups):
                temp_index = 0
            print("picked up", temp_cups[temp_index])
            removed_cups.append(temp_cups[temp_index])
            del temp_cups[temp_index]
        while destination_cup not in temp_cups:
            destination_cup = (destination_cup - 1) % 10
        print("destination cup", destination_cup)
        destination_cup_index = temp_cups.index(destination_cup)
        for i, cup in enumerate(removed_cups):
            temp_cups.insert(destination_cup_index+1+i, cup)
        current_cup_index = (temp_cups.index(
            current_cup)+1) % len(temp_cups)
        current_cup = temp_cups[current_cup_index]
    result = ""
    start_index = temp_cups.index(1)+1
    for i in range(len(temp_cups)-1):
        result += str(temp_cups[(start_index+i) % len(temp_cups)])
    print(result)
    return result


class ListNode:
    def __init__(self, value: int):
        self.value = value
        self.next: 'ListNode' = None  # type: ignore

    def add(self, node: 'ListNode'):
        cursor = self
        while cursor.next:
            cursor = cursor.next
        cursor.next = node

    def __eq__(self, other):
        return self.value == other.value


def print_ll(head):
    cursor = head
    while cursor.value != 1:
        cursor = cursor.next
    cursor = cursor.next
    while True:
        print(cursor.value, end="")
        cursor = cursor.next
        if cursor.value == 1:
            break
    print()


def solve(numbers: list[int], iterations: int) -> ListNode:
    max_number = sorted(numbers)[-1]

    map_of_nodes = {}
    head = ListNode(numbers[0])
    map_of_nodes[numbers[0]] = head
    cursor = head
    for number in numbers[1:]:
        cursor.next = ListNode(number,)
        cursor = cursor.next
        map_of_nodes[cursor.value] = cursor
    cursor.next = head
    cursor = head
    for i in range(iterations):
        if i % 5000 == 0:
            print(f"{i}/{iterations}\r", end="")
        val = cursor.value - 1
        if val <= 0:
            val = max_number

        next_vals = []
        next_vals.append(cursor.next.value)
        next_vals.append(cursor.next.next.value)
        next_vals.append(cursor.next.next.next.value)

        while val in next_vals:
            val -= 1
            if val <= 0:
                val = max_number

        three_cups = cursor.next  # get three cups
        cursor.next = cursor.next.next.next.next  # remove them from the list

        tmp_cursor = map_of_nodes[val]

        tmp_next = tmp_cursor.next
        tmp_cursor.next = three_cups
        tmp_cursor.next.next.next.next = tmp_next
        cursor = cursor.next

    print()
    return head


def part2(lines):
    numbers = [int(x) for x in lines[0]]
    # print("Part 1:")
    # print_ll(solve(numbers, 100))
    numbers.extend(list(range(10, 1000001)))
    llist = solve(numbers, 10_000_000)
    cursor = llist
    while cursor.value != 1:
        cursor = cursor.next
    return cursor.next.value * cursor.next.next.value
    
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
