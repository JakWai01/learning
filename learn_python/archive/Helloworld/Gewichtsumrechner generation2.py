weight = int(input("please enter your weight: "))
converter = input("to (k)g or to (l)bs: ")


if converter.upper() == "K":
    result = weight * 0.454
    print(f"Your weight is {result}")
if converter.upper() == "L":
    result = weight * 2.2046
    print(f"Your weight is {result}")
