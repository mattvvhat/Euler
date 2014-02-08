#include <iostream>
#include "lib/BigInteger.h"

BigInteger C (unsigned n, unsigned r) {
  BigInteger mul(1);

  for (int i=2; i <= n; i++) {
    mul *= i;
  }

  for (int i=2; i <= r; i++) {
    mul /= i;
  }

  for (int i=2; i <= n-r; i++) {
    mul /= i;
  }

  return mul;
}

int main () {

  unsigned count=0;

  for (unsigned n=1; n <= 100; n++) {
    for (unsigned r=2; r <= n; r++) {
      BigInteger what = C(n, r);
      if (what.getNumber().length() > 6) {
        count++;
      }
    }
  }

  std::cout << "\n" << count << "\n";

  return 0;
}
