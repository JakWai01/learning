# Das hier ist die Liste
numbers = [1, 2, 3, 4, 5, 6]
#Aufgabe ist es, die größte Zahl in der Liste zu finden

max = numbers[0]

for number in numbers:
    if number > max:
        max = number
print(max)



