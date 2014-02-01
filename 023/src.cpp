#include <iostream>
#include <fstream>
#include "aljebr.h"
#include <set>

bool is_abundant (int n) {
  int val = al::divisor_sum(n);
  return val > n;
}

#define MAX 28123

int main () {

  std::ofstream out;
  out.open("abundant_numbers.txt");

  std::vector <int> abundants;
  std::set <int> sums;

  for (int i=0; i < MAX; i++) {
    if (is_abundant(i))
      out << i << "\n";
  }

  out.close();

  return 0;
}
