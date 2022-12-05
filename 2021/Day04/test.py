seq = []
cards = []

with open("data/test_input.txt") as file:
    seq = next(file).split(",")

    card = []
    for row in file:
        row = row.strip()
        if not row:
            if card:
                cards.append(card)
            card = []
            continue
        card.append(row.split())
    cards.append(card)

print(seq)
print(cards)