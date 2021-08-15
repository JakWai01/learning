row1 = [" ", " ", " "]
row2 = [" ", " ", " "]
row3 = [" ", " ", " "]
game = True
winner = False
count = 0
def printGame():
    print(row1)
    print(row2)
    print(row3) 

def checkWin(): 

    if row1[0] == "X" and row1[1] == "X" and row1[2] == "X":
        print("Player 1 won!")
        return True
    if row2[0] == "X" and row2[1] == "X" and row2[2] == "X":
        print("Player 1 won!")
        return True
    if row3[0] == "X" and row3[1] == "X" and row3[2] == "X":
        print("Player 1 won!")
        return True
    if row1[0] == "X" and row2[0] == "X" and row3[0] == "X":
        print("Player 1 won!")
        return True
    if row1[1] == "X" and row2[1] == "X" and row3[1] == "X":
        print("Player 1 won!")
        return True
    if row1[2] == "X" and row2[2] == "X" and row3[2] == "X":
        print("Player 1 won!")
        return True
    if row1[0] == "X" and row2[1] == "X" and row3[2] == "X":
        print("Player 1 won!")
        return True
    if row1[2] == "X" and row2[1] == "X" and row3[0] == "X":
        print("Player 1 won!")  
        return True
    if row1[0] == "O" and row1[1] == "O" and row1[2] == "O":
        print("Player 2 won!")
        return True
    if row2[0] == "O" and row2[1] == "O" and row2[2] == "O":
        print("Player 2 won!")
        return True
    if row3[0] == "O" and row3[1] == "O" and row3[2] == "O":
        print("Player 2 won!")
        return True
    if row1[0] == "O" and row2[0] == "O" and row3[0] == "O":
        print("Player 2 won!")
        return True
    if row1[1] == "O" and row2[1] == "O" and row3[1] == "O":
        print("Player 2 won!")
        return True
    if row1[2] == "O" and row2[2] == "O" and row3[2] == "O":
        print("Player 2 won!")
        return True
    if row1[0] == "O" and row2[1] == "O" and row3[2] == "O":
        print("Player 2 won!")
        return True
    if row1[2] == "O" and row2[1] == "O" and row3[0] == "O":
        print("Player 2 won!")
        return True
    if row1[0] != " " and row1[1] != " " and row1[2] != " " and row2[0] != " " and row2[1] != " " and row2[2] != " " and row3[0] != " " and row3[1] != " " and row3[2] != " ":
        print("Draw!")
        return True

while game == True: 

    printGame()

    while winner == False:

        if count % 2 == 0:
            print("Player 1: Please enter where you want to make your cross [1-9]")
            num = input()
            
            if num == "1": 
                row1[0] = "X"
            if num == "2":
                row1[1] = "X"
            if num == "3":
                row1[2] = "X"
            if num == "4":
                row2[0] = "X"
            if num == "5":
                row2[1] = "X"
            if num == "6":
                row2[2] = "X"
            if num == "7":
                row3[0] = "X"
            if num == "8":
                row3[1] = "X"
            if num == "9":
                row3[2] = "X"

        if count % 2 != 0:
            print("Player 2: Please enter where you want to make your circle [1-9]")
            num = input()

            if num == "1": 
                row1[0] = "O"
            if num == "2":
                row1[1] = "O"
            if num == "3":
                row1[2] = "O"
            if num == "4":
                row2[0] = "O"
            if num == "5":
                row2[1] = "O"
            if num == "6":
                row2[2] = "O"
            if num == "7":
                row3[0] = "O"
            if num == "8":
                row3[1] = "O"
            if num == "9":
                row3[2] = "O"

        count += 1
        printGame()  
        if checkWin():
            winner = True

    count = 0
    winner = False
    row1 = [" ", " ", " "]
    row2 = [" ", " ", " "]
    row3 = [" ", " ", " "]

    print("Another round? 1 = yes 2 = no")

    choice = input()

    if choice == "2": 
        game = False