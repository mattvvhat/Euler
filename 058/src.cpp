#include <iostream>
#include "aljebr.h"

int main () {

  auto br = [&] (int n) -> int { return (2*n+1)*(2*n+1); };
  auto bl = [&] (int n) -> int { return (2*n+1)*(2*n+1) - 2*n; };
  auto ul = [&] (int n) -> int { return (2*n+1)*(2*n+1) - 4*n; };
  auto ur = [&] (int n) -> int { return (2*n+1)*(2*n+1) - 6*n; };

  unsigned n=1;
  unsigned number_primes=0;
  double ratio;
  unsigned side_length=1;
  unsigned k=1;

  do {
    side_length += 2;

    if (al::is_prime(br(k)))
      number_primes++;
    if (al::is_prime(bl(k)))
      number_primes++;
    if (al::is_prime(ul(k)))
      number_primes++;
    if (al::is_prime(ur(k)))
      number_primes++;

    ratio = (float) number_primes / (1 + 4*k);

    std::cout << "ratio = " << number_primes << "/" << (1+4*k) << " = " << ratio << "\n";

    k++;
  } while (ratio > .1);

  std::cout << side_length << "\n";

  return 0;
}
