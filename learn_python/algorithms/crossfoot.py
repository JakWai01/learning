number = int(input())

def crossfoot(x: int) -> int:
    if x < 0:
        return crossfoot(-x)
    else:
        if x < 10:
            return x
        else:
            return crossfoot(x//10) + (x % 10)

print(crossfoot(number))
