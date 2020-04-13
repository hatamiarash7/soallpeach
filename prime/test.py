import math, sys

@profile
def is_prime(n: int) -> bool:
    for x in range(2, int(math.sqrt(n))):
        if n % x == 0:
            return False
    else:
        return True

with open('input.txt') as numbers:
    for number in numbers:
        print(1 if is_prime(int(number)) else 0)
