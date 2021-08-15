def make_chocolate(small, big, goal):
    # Wenn die Gesamtschokolade nicht ausreicht
    if small + big * 5 < goal:
        return -1

    # Wenn small kleiner ist als der Rest:
    if goal % 5 > small:
        return -1

    # Wenn der Rest != 0 ist, aber schon vorher nicht genug big vorhanden waren
    if goal / 5 > big:
        result0 = goal - big * 5
        return result0

    # Wenn der Rest 0 ist, aber nicht genug big da sind
    if goal % 5 == 0 and goal / 5 > big:
        result1 = (goal / 5 - big) * 5
        if small > result1:
            return result1

    # Wenn small größer ist als der Rest:
    if goal % 5 < small:
        return goal % 5
    # Jeder andere Fall wenn das Result einfach nur der Rest ist und keine andere condition applied
    else:
        result = goal % 5
        return result
