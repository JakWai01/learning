digits = [0,1,2,3,4,5,6,7,8,9]
#funktioniert auch mit strings
print(digits[::-2])

for i in range(len(digits)):
    print(digits[0:i])

for i in range(len(digits)):
    print(digits[i:i+3])