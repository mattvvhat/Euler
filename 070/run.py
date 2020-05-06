#!/usr/bin/env python


from collections import defaultdict


def gcd(a, b):

    while a != b:
        if a < b:
            a, b = b, a
        a, b = a-b, b

    return a


def is_permutation(a, b):
    a = str(a)
    b = str(b)

    if len(a) != len(b):
        return False

    
    for lhs, rhs in zip(sorted(a), sorted(b)):
        if lhs != rhs:
            return False

    return True


def phi(n):
    hits = 0
    for i in range(2, n):
        if n % i != 0 and gcd(n, i) == 1:
            hits += 1
    return hits + 1


def tot(n):
    nom = 1
    den = 1
    for p in primes(n).keys():
        nom *= p-1
        den *= p
    return n / den * nom


def sieve(n):
    tracker = defaultdict(lambda: True)

    for i in range(2, n):

        if not tracker[i]:
            continue

        j = i*i

        while j <= n:
            tracker[j] = False
            j += i

    return sorted([
        i
        for i in range(2, n)
        if tracker[i]
    ])


PRIMES = sieve(10**7)
PRIME_SET = set(PRIMES)


def primes(n):
    factors = defaultdict(int)

    i = 0
    v = n

    while v != 1:

        # Early exit if this is prime
        if v in PRIME_SET:
            factors[v] = 1
            break

        d = PRIMES[i]

        while v % d == 0:
            factors[d] += 1
            v /= d

        i += 1

    return factors


if __name__ == "__main__":

    vals = [
        i
        for i in range(2, 10**7)
        if is_permutation(i, tot(i))
    ]

    vals = sorted(vals, key=lambda n: n/float(tot(n)))

    print vals[0]
