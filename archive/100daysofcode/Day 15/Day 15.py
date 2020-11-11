from random import randint


t = ["Rock", "Paper", "Scissors"]
AGAIN = 1
while AGAIN == 1:
    EINGABE = input("Rock, Paper, Scissors?")
    bot = t[randint(0, 2)]

    if EINGABE == bot:
        print("Tie!")
    elif EINGABE == "Rock":
        if bot == "Paper":
            print("You loose!")
        else:
            print("You win!")
    elif EINGABE == "Paper":
        if bot == "Rock":
            print("You win!")
        else:
            print("You loose!")
    elif EINGABE == "Scissors":
        if bot == "Rock":
            print("You loose!")
        else:
            print("You win!")
    print("Do you want to play again? [Y][N]")
    repeat = input()
    if repeat == "Y":
        AGAIN = 1
        continue
    elif repeat == "N":
        AGAIN = 0
        break