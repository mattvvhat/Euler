#include <iostream>
#include "aljebr.h"
#include "sieve.hpp"

int main () {
  Eratosthenes::expand(1000000);

  std::cout << "------\n";

  for (int i=0; i < 1000000; i++) {
    if (!Eratosthenes::is_prime(i)) {
      if (al::is_prime(i)) {
        std::cout << "ISSUE (i=" << i << ")\n";
      }
    }
  }

  return 0;
}
