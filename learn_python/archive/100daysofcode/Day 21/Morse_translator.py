print("ENTER TO TRANSLATE(letters separated by a space):")
translate = input().split()

for element in translate:
    if element == ".-":
        print("a", end="")
    elif element == "-...":
        print("b", end="")
    elif element == "-.-.":
        print("c", end="")
    elif element == "-..":
        print("d", end="")
    elif element == ".":
        print("e", end="")
    elif element == "..-.":
        print("f", end="")
    elif element == "--.":
        print("g", end="")
    elif element == "....":
        print("h", end="")
    elif element == "..":
        print("i", end="")
    elif element == ".---":
        print("j", end="")
    elif element == "-.-":
        print("k", end="")
    elif element == "-.--":
        print("l", end="")
    elif element == "--":
        print("m", end="")
    elif element == "-.":
        print("n", end="")
    elif element == "---":
        print("o", end="")
    elif element == ".--.":
        print("p", end="")
    elif element == "--.-":
        print("q", end="")
    elif element == ".-.":
        print("r", end="")
    elif element == "...":
        print("s", end="")
    elif element == "-":
        print("t", end="")
    elif element == "..-":
        print("u", end="")
    elif element == "...-":
        print("v", end="")
    elif element == ".--":
        print("w", end="")
    elif element == "-..-":
        print("x", end="")
    elif element == "-.--":
        print("y", end="")
    elif element == "--..":
        print("z", end="")
    elif element == "-----":
        print("0", end="")
    elif element == ".----":
        print("1", end=", ")
    elif element == "..---":
        print("2", end="")
    elif element == "...--":
        print("3", end="")
    elif element == "....-":
        print("4", end="")
    elif element == ".....":
        print("5", end="")
    elif element == "-....":
        print("6", end="")
    elif element == "--...":
        print("7", end="")
    elif element == "---..":
        print("8", end="")
    elif element == "----.":
        print("9", end="")
    elif element == "a":
        print(".- ", end=" ")
    elif element == "b":
        print("-... ", end=" ")
    elif element == "c":
        print("-.-. ", end=" ")
    elif element == "d":
        print("-.. ", end=" ")
    elif element == "e":
        print(". ", end=" ")
    elif element == "f":
        print("..-. ", end=" ")
    elif element == "g":
        print("--. ", end=" ")
    elif element == "h":
        print(".... ", end=" ")
    elif element == "i":
        print(".. ", end=" ")
    elif element == "j":
        print(".--- ", end=" ")
    elif element == "k":
        print("-.- ", end=" ")
    elif element == "l":
        print("-.-- ", end=" ")
    elif element == "m":
        print("-- ", end=" ")
    elif element == "n":
        print("-. ", end=" ")
    elif element == "o":
        print("--- ", end=" ")
    elif element == "p":
        print(".--. ", end=" ")
    elif element == "q":
        print("--.- ", end=" ")
    elif element == "r":
        print(".-. ", end=" ")
    elif element == "s":
        print("... ", end=" ")
    elif element == "t":
        print("- ", end=" ")
    elif element == "u":
        print("..- ", end=" ")
    elif element == "v":
        print("...- ", end=" ")
    elif element == "w":
        print(".-- ", end=" ")
    elif element == "x":
        print("-..- ", end=" ")
    elif element == "y":
        print("-.-- ", end=" ")
    elif element == "z":
        print("--.. ", end=" ")
    elif element == "0":
        print("----- ", end=" ")
    elif element == "1":
        print(".---- ", end=" ")
    elif element == "2":
        print("..--- ", end=" ")
    elif element == "3":
        print("...-- ", end=" ")
    elif element == "4":
        print("....- ", end=" ")
    elif element == "5":
        print("..... ", end=" ")
    elif element == "6":
        print("-.... ", end=" ")
    elif element == "7":
        print("--... ", end=" ")
    elif element == "8":
        print("---.. ", end=" ")
    elif element == "9":
        print("----. ", end=" ")
    elif element == " ":
        print(" ", end=" ")
