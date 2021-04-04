
class Person:
    number_of_people = 0

    def __init__(self, name):
        self.name = name

    @classmethod
    def num(cls):
        return cls.number_of_people
    
    @staticmethod 
    def add5(x):
        return x + 5
p1 = Person("tim")
p2 = Person("jill")

print(p1.number_of_people)
print(p2.number_of_people)
Person.number_of_people = 1
print(p1.number_of_people)
print(p2.number_of_people)
print(Person.number_of_people)
print(Person.num())
print(Person.add5(4))
