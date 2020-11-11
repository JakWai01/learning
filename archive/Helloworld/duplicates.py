numbers = [1, 2, 4, 4, 6, 7]
singles = []

for number in numbers:
    if number not in singles:
        singles.append(number)
print(singles)
singles.reverse()
print(singles)
