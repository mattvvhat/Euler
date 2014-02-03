#include <iostream>
#include <map>
#include "Eratosthenes.hpp"

#define MAX 200000

typedef unsigned long long number;

std::map <number, number> factors (number n) {
  std::map <number, number> divs;
  auto it   = Eratosthenes::begin();
  auto end  = Eratosthenes::end();
  while (it != end) {
    number d = *it;
    while (n % d == 0) {
      n /= d;
      divs[d] = divs.count(d) + 1;
    }
    ++it;
  }
  return divs;
}

int main () {

  // NOTE: It seems like the sieve is being made too slowly...
  Eratosthenes::expand(MAX);

  std::cout << "Sieve created";

  int i, n1, n2, n3, n4;
  n1 = n2 = n3 = 0;
  for (i=100000; i < MAX; i++) {
    n4 = factors(i).size();

    if (n1 == n2 && n2 == n3 && n3 == n4 && n4 == 4) {
      std::cout << i << "\n";
      break;
    }

    n1 = n2;
    n2 = n3;
    n3 = n4;
  }

  std::cout << i-3 << "(" << factors(i-3).size() << ")\n";
  std::cout << i-2 << "(" << factors(i-2).size() << ")\n";
  std::cout << i-1 << "(" << factors(i-1).size() << ")\n";
  std::cout << i-0 << "(" << factors(i-0).size() << ")\n";

  return 0;
}
