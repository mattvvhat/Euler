#include <iostream>

#include "aljebr.h"
#include "Eratosthenes.hpp"

#define SIZE 1000000

int main () {

  unsigned long largest=0;
  unsigned long val;
  Eratosthenes::expand(SIZE);

  for (auto start=Eratosthenes::begin(); start != Eratosthenes::end(); ++start) {

    unsigned count=1;
    unsigned long sum=0;

    for (auto it=start; it != Eratosthenes::end(); ++it) {
      sum += *it;

      if (sum >= SIZE)
        break;

      if (count > largest && al::is_prime(sum)) {
        largest = count;
        std::cout << sum << " (" << count << ") ~ " << al::is_prime(sum) << "\n";
      }

      count++;
    }

  }

  std::cout << largest << "\n";

  return 0;
}
