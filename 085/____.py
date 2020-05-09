#!/usr/bin/env python3


import math
from collections import defaultdict



class Grid(object):
    def __init__(self, pos, dim):
        if dim[0] == 0 or dim[1] == 0:
            raise "Wtf"
        self.pos = pos
        self.dim = dim

    def __contains__(self, grid):

        if  grid.pos[0] < self.pos[0]:
            return False

        if grid.pos[0]+grid.dim[0] > self.pos[0]+self.dim[0]:
            return False

        if grid.pos[1] < self.pos[1]:
            return False

        if grid.pos[1]+grid.dim[1] > self.pos[1]+self.dim[1]:
            return False

        return True
    
    def key(self):
        return (self.pos[0],  self.pos[1], self.dim[0], self.dim[1])

    def __repr__(self):
        return f"Grid({self.pos[0]}, {self.pos[1]}, {self.dim[0]}x{self.dim[1]})"


CACHE = defaultdict(int)

def count_rectangles(grid):
    """Naive counting algorithm"""
    count = 0
    for i in range(grid.dim[0]):
        for j in range(grid.dim[1]):
            x = grid.pos[0] + i
            y = grid.pos[1] + j
            count += (grid.dim[0]-x)*(grid.dim[1]-y)
    return count


def compute(x, y, w, h):
    assert x >= 0 and y >= 0 and w > 0 and h > 0
    if w == 1 and h == 1:
        return 1



def buildup(width, height):
    results = {(1, 1, 1, 1): 1}


    count = defaultdict(int)

    for i in range(1, width+1):
        for j in range(1, height+1):

            for y in range(0, width-i+1):
                for x in range(0, height-j+1):
                    count[i+y, j+x] += i*j



    return count




if __name__ == "__main__":


    SEARCH_HEIGHT = 90
    SEARCH_WIDTH = 90

    vals = buildup(SEARCH_HEIGHT, SEARCH_WIDTH)


    TARGET = 2000000
    BEST_DISTANCE = TARGET
    BEST_VALUE = -1
    BEST_INDEX = (-1, -1)

    for (y, x), count in vals.items():

        distance = abs(TARGET - count)

        if distance <= BEST_DISTANCE:
            BEST_DISTANCE = distance
            BEST_VALUE = count
            BEST_INDEX = (y, x)



    print("DISTANCE ..", BEST_DISTANCE)
    print("VALUE .....", BEST_VALUE)
    print("INDEX .....", BEST_INDEX)
    print("...........", vals.get(BEST_INDEX))
    print()
    print("AREA =", BEST_INDEX[0]*BEST_INDEX[1])
