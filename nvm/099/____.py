#!/usr/bin/env python3


import math
from collections import namedtuple


ExpPair = namedtuple("ExpPair", ["base", "exp"])


def change_base(e, base):
    return ExpPair(
        base,
        e.exp * math.log(e.base, base),
    )


class ExpPair(object):
    def __init__(self, base, exp):
        self.base = base
        self.exp = exp

    def value(self):
        return self.base**self.exp
    
    def __lt__(self, rhs):
        changed = change_base(rhs, self.base)
        return self.exp < changed.exp
    
    def __repr__(self):
        return f"{self.base}**{self.exp}"


if __name__ == "__main__":
    e = ExpPair(2, 4)
    b = change_base(e, 4)

    e = ExpPair(17, 23)
    b = change_base(e, 4)
    # print(f(e))
    # print(f(b))
    print(ExpPair(2, 11) < ExpPair(3, 7))
    print(ExpPair(632382, 518061) < ExpPair(519432, 525806))


    lines = []

    with open("p099_base_exp.txt", "r") as fi:

        line_number = 0

        largest = ExpPair(1.01, 1)
        largest_index = -1

        for line in fi:

            line_number += 1

            b, p = line.strip().split(",", 2)

            v = ExpPair(
                int(b),
                int(p),
            )

            if largest < v:
                largest = v
                largest_index = line_number

    print(largest_index, "=>", largest)
