#include <iostream>
#include "Eratosthenes.hpp"

#define MAX 10000

bool has_goldbach (unsigned n) {
  for (unsigned i=0; i <= n; i+=1) {
    if (Eratosthenes::is_prime(n-2*i*i)) {
      return true;
    }
  }
  return false;
}

int main () {

  Eratosthenes::expand(MAX);

  for (unsigned i=2; i < MAX; i++) {
    unsigned val = 2*i+1;
    if (!Eratosthenes::is_prime(val)) {

      if (!has_goldbach(val)) {
        std::cout << val << "\n";
        break;
      }
      // std::cout << val << "\n";
    }
  }

  return 0;
}
