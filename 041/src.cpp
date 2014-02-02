#include <iostream>
#include <algorithm>
#include <string>

#include "aljebr.h"

int main () {

  unsigned long largest = 0;
  std::string number = "1234567";

  do {
    unsigned long n = atoi(number.c_str());
    if (al::is_prime(n) && n > largest) {
      largest = n;
    }
  } while (std::next_permutation(number.begin(), number.end()));

  std::cout << largest << "\n";

  return 0;
}
