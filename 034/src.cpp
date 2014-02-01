#include <iostream>
#include "lib/BigInteger.h"
#include <string>
#include <cmath>

typedef unsigned long long int ull;

ull factorial (unsigned n) {
  ull m = 1;
  while (n > 1)
    m *= n--;
  return m;
}

ull factorial_sum (unsigned n) {
  std::string str;
  char rawstr[256];
  sprintf(rawstr, "%u", n);
  str = rawstr;
  unsigned sum = 0;
  for (std::string::iterator it=str.begin(); it != str.end(); ++it) {
    unsigned c = *it - '0';
    sum += factorial(c);
  }
  return sum;
}

int main () {

  // 
  ull max = pow(10, 7);

  unsigned total = 0;

  for (unsigned i=10; i < max; i++) {
    if (i == factorial_sum(i)) {
      total += i;
    }
  }

  // 
  std::cout << total << "\n";

  // 
  return 0;
}
