from typing import Callable
from functools import reduce

def multiply_list(xs: list, k: int) -> list:
    return [k * i for i in xs]
print(multiply_list([0,2,3,4,5], 3))

# Closures 
def multiply(x: int) -> Callable[[int], int]:
    def multiply_x(y: int) -> int:
        return x * y
    return multiply_x
mal2 = multiply(2)
print(mal2(2))

def multiply_list_functional(xs: list, k: int) -> list: 
    return list(map(lambda x: x * k, xs))
print(multiply_list_functional([3,4,5], 5))

def bigger_than(xs: list, k: int) -> list:
    return [x for x in xs if x > k]
print(bigger_than([1,3,3,4,5], 3))

def summe(xs: list) -> int:
    return reduce(lambda x1, x2: x1 + x2, xs, 0)
print(summe([1,2,3,4,5,6,7,8,9]))
    
