t = [1, 2, 3]

credit_card1 = (12312412312312312, "Jakob Waibel", "11/20", 123)
credit_card2 = (12312412312312312, "Jakob Waibel", "11/20", 123)

credit_cards = [credit_card1, credit_card2]

print(credit_cards)

person1 = ("nancy-pants", 25, "Pizza")
person2 = ("sdfsd-pants", 25, "Pizza")
people = [person1, person2]
#(name, age, favfood) = person

#print(name)
#print(age)
#print(favfood)

for name, age, favfood in people:
    print(name)
    print(age)
    print(favfood)
    print()