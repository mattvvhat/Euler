#include <iostream>
#include "lib/BigInteger.h"



int main () {
  BigInteger y0(1), y1(1), y2;
  BigInteger n;

  int k=2;

  while (y2.getNumber().length() < 1000) {
    y2 = y0 + y1;
    y0 = y1;
    y1 = y2;

    n = y2;
    k++;
  }

  std::cout << n.getNumber() << " (" << k << ")\n";

  return 0;
}
