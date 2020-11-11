guesses = 1

while guesses <=3:
    number = input("Enter the number: ")
    guesses = guesses + 1
    if number == "9":
        print("Damn right!")
        break

if guesses == 4:
    print("No guesses Left")