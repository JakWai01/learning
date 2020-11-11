#lists
#es können natürlich auch strings in lists sein
l = [1, 2, 3]
l2 = [1, "string", 3, True, [1, 2, 3]]
print(l)
print(l2)

print(l2[0])
print(l2[1])
print(l2[2])
print(l2[3])
print(l2[4])

names = ["joe", "john", "james"]

names.append("gary")
names.insert(0, "firsty")
print(names)
names.remove("john")
print(names)

numbers = [4, 2, 3, 4]

print(numbers)
numbers.sort()
print(numbers)
for number in numbers:
    print(number)