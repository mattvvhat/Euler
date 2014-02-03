#include <iostream>
#include "Eratosthenes.hpp"
#include "aljebr.h"

unsigned consecutive_prime_count (unsigned a, unsigned b) {

  // NOTE: This loop goes on indefinitely
  for (unsigned k=0; true; k++) {
    if (!al::is_prime(k*k + a*k + b))
      return k;
  }

  // This should never happen
  return 0;
}


int main () {

  Eratosthenes::expand(1000);

  std::cout << "Sieve created\n";

  struct {
    unsigned count = 0;
    int a, b;
  } largest;

  for (auto it=Eratosthenes::begin(); it != Eratosthenes::end(); ++it) {
    unsigned b = *it;

    for (int a=-1000; a <= +1000; a++) {
      if (al::is_prime(1 + a + b)) {
        int count = consecutive_prime_count(a, b);
        if (count > largest.count) {
          largest.count = count;
          largest.a = a;
          largest.b = b;
          std::cout << largest.a * largest.b << "\n";
        }
      }
    }
  }

  std::cout << largest.a << " " << largest.b << "\n";
  std::cout << largest.a * largest.b << "\n";

  return 0;
}
