from collections import OrderedDict


dictionary = {"Jakob" : 18, "Leon" : 19}

print(dictionary["Jakob"])

contacts = {
    "Joe" : [1234567 , "emailadress@website.com"],
    "Jane" : [2345678, "work@web.de"]
}

print(contacts["Joe"])
sentence = "I like the name Jakob because the name Jakob is the best"

word_counts = {
    "I": 1,
    "like": 1,
    "the": 3
}

 #dict.items()
 #dict.keys()
 #dict.values()

 #word_counts.pop("I")
# print(word_counts)

print(word_counts)
#word_counts["Aaron"] = 2 adding to dictionary
print(word_counts.popitem())
print(word_counts)

#dict-clear()
#ord_counts.clear()
print(word_counts)

print(list(word_counts.keys()))
#Ich liebe es. Man kann hier einfach key oder value eingeben


