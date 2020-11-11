# list comprehentions
name = ["Jennifer", "Susan", "Mike", "Jakob"]

l = []
for person in name:
    l.append(person)
print(l)
print([person + " dumped me" for person in name])

l = []
for person in name:
    l.append(person + " dumped me.")
print(l)

movies_and_rating = {"Interstellar": 9, "Dark Knight": 8, "50 Shades": 3}

p = []

for movie in movies_and_rating:
    if movies_and_rating[movie] > 6:
        p.append(movie)

print(p)

print( [movie for movie in movies_and_rating if movies_and_rating[movie] > 6] )

