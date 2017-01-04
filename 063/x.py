import math

s = 0
b = 1

while True:

    new_sum = 0

    for a in range(1, 10):
        v = a**b
        digits = len(str(v))

        print v, a, b, digits
        if b == digits:
            new_sum += 1

    s += new_sum

    print new_sum

    if new_sum == 0:
        break
    else:
        b += 1

print s
