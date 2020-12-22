import sys


def part1(lines):
    player_1_cards = []
    player_2_cards = []
    i = 1
    while len(lines[i]) != 0:
        player_1_cards.append(int(lines[i]))
        i += 1
    i += 2
    while i < len(lines) and len(lines[i]) != 0:
        player_2_cards.append(int(lines[i]))
        i += 1
    while len(player_1_cards) > 0 and len(player_2_cards) > 0:
        player_1_play = player_1_cards.pop(0)
        player_2_play = player_2_cards.pop(0)
        if player_1_play > player_2_play:
            player_1_cards.extend([player_1_play, player_2_play])
        else:
            player_2_cards.extend([player_2_play, player_1_play])
    result = 0
    for multiplier, card in enumerate((player_1_cards + player_2_cards)[::-1]):
        result += card * (multiplier + 1)
    return result


def recursive_game(player_1_cards, player_2_cards, depth=0):
    history = set()
    while len(player_1_cards) > 0 and len(player_2_cards) > 0:
        winner = 0
        # player 1 wins if this exact combination exists in history
        if tuple(player_1_cards+[-1]+player_2_cards) in history:
            return 1, player_1_cards
        history.add(tuple(player_1_cards+[-1]+player_2_cards))
        player_1_play = player_1_cards.pop(0)
        player_2_play = player_2_cards.pop(0)

        if player_1_play <= len(player_1_cards) and player_2_play <= len(player_2_cards):
            winner, _ = recursive_game(
                list(player_1_cards[:player_1_play]), list(player_2_cards[:player_2_play]), depth=depth+1)
        else:
            if player_1_play > player_2_play:
                winner = 1
            else:
                winner = 2
        if winner == 1:
            player_1_cards.extend([player_1_play, player_2_play])
        else:
            player_2_cards.extend([player_2_play, player_1_play])

    if len(player_1_cards) == 0:
        return 2, player_2_cards
    else:
        return 1, player_1_cards


def part2(lines):
    player_1_cards = []
    player_2_cards = []
    i = 1
    while len(lines[i]) != 0:
        player_1_cards.append(int(lines[i]))
        i += 1

    i += 2
    while i < len(lines) and len(lines[i]) != 0:
        player_2_cards.append(int(lines[i]))
        i += 1
    _, winner_deck = recursive_game(player_1_cards, player_2_cards)
    result = 0
    for multiplier, card in enumerate((winner_deck)[::-1]):
        result += card * (multiplier + 1)
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
