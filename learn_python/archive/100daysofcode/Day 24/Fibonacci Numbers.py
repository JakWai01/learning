num1 = 1
num2 = 1
print("PLEASE ENTER A NUMBER: ")
n = int(input())
value = 3

if n == 1:
    print("1")
if n == 2:
    print("1")
    print("1")
elif n >= 2:
    print("1")
    print("1")
    while value <= n:
        result = num1 + num2
        num1 = num2
        num2 = result
        print(result)
        value += 1
