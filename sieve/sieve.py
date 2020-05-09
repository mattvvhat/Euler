#!/usr/bin/env python


from collections import defaultdict



def sieve(n):
    tracker = defaultdict(lambda: True)

    for i in range(2, n):

        if not tracker[i]:
            continue

        j = i*i

        while j <= n:
            tracker[j] = False
            j += i

    return [
        i
        for i in range(2, n)
        if tracker[i]
    ]


if __name__ == "__main__":
    PRIMES = sieve(10**7)
    print "COUNT =>", len(PRIMES)
