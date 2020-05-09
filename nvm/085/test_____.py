#!/usr/bin/env python3

import unittest

from ____ import Grid
from ____ import count_rectangles
from ____ import compute
from ____ import buildup


class TestGrid(unittest.TestCase):
    def test_contains(self):
        cases = [
            Grid([0, 0], [1, 1]),
            Grid([0, 0], [2, 2]),
            Grid([0, 0], [3, 3]),
        ]

        for i in range(len(cases)-1):
            self.assertTrue(cases[i] in cases[i+1])

        self.assertNotIn(
            Grid([0, 0], [2, 1]),
            Grid([1, 1], [1, 1]),
        )

        self.assertIn(
            Grid([1, 0], [1, 1]),
            Grid([0, 0], [2, 2]),
        )

        self.assertNotIn(
            Grid([2, 0], [1, 1]),
            Grid([0, 0], [2, 2]),
        )

        # Y-coordinates

        self.assertNotIn(
            Grid([0, 0], [1, 1]),
            Grid([0, 1], [2, 2]),
        )

        self.assertNotIn(
            Grid([1, 1], [1, 4]),
            Grid([0, 0], [2, 2]),
        )

        self.assertIn(
            Grid([1, 1], [1, 1]),
            Grid([0, 0], [2, 2]),
        )

        self.assertNotIn(
            Grid([2, 2], [1, 1]),
            Grid([0, 0], [2, 2]),
        )


class TestRectCount(unittest.TestCase):

    def test_simple(self):

        cases = [
            (Grid([0, 0], [1, 1]), 1),
            (Grid([0, 0], [1, 2]), 3),
            (Grid([0, 0], [2, 1]), 3),
            (Grid([0, 0], [2, 2]), 9),
            (Grid([0, 0], [3, 2]), 18),
        ]

        for grid, count in cases:
            self.assertEqual(count_rectangles(grid), count, f"Count({grid}) != {count}")


    def test_okay(self):

        results = buildup(2, 2)
        self.assertEqual(len(results), 2*2)

        HEIGHT = 70
        WIDTH = 70
        results = buildup(HEIGHT, WIDTH)
        print(results.get((HEIGHT, WIDTH)))

    @unittest.skip("x_x")
    def test_recurse(self):

        cases = {
            (0, 0, 1, 1): 1,
            (1, 1, 1, 1): 1,
            (2, 2, 1, 1): 1,
            (0, 0, 2, 1): 3,
            (0, 0, 1, 2): 3,
        }

        for key, count in cases.items():
            self.assertEqual(compute(*key), count, f"{key} != {count}")


if __name__ == "__main__":
    unittest.main()
