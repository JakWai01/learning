from typing import Callable


def multipliziere_liste(xs: list, k: int) -> list:
    ausgabe = []
    for x in xs:
        ausgabe.append(x * k)
    return ausgabe


def mal(x: int) -> Callable[[int], int]:
    def mal_x(y: int) -> int:
        return x * y

    return mal_x

# mal_drei = mal(3)
# print(mal_drei(5))

def multipliziere_liste_funktional(xs: list, k: int) -> list:
    return list(map(lambda x: x * k, xs))


# print(multipliziere_liste_funktional([1,2,3], 3))

def groesser_als(xs: list, k: int) -> list:
    return list(filter(lambda x: x > k, xs))


# print(groesser_als([1,2,3,4,5], 3))

def multipliziere_liste_funktional_python(xs: list, k: int) -> list:
    return [x * k for x in xs]


def groesser_als_python(xs: list, k: int) -> list:
    return [x for x in xs if x > k]
