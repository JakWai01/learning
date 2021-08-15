#sets (Mengen)
#code übersichtlicher gestalten
l =[1, 2, 3, 4, 4, 6, 6, 6]
print(l)
q = {1, 1, 1, 1, 1, 1}
print(q)
s = {"banana", "apple"}
s.add("ananas")
s.add(4)
print(s)
a = set(l)
print(a)
library1 = {1,2,3}
library2 = {3,4,5}

library = library1.union(library2)

print(library)

library3 = library1.difference(library2)
print(library3)