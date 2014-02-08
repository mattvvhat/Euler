#include <iostream>
#include <string>
#include "lib/BigInteger.h"

BigInteger pow (unsigned a, unsigned b) {
  BigInteger m(1);
  while (b-- > 0)
    m *= a;
  return m;
}

unsigned digit_sum (BigInteger n) {
  unsigned sum=0;
  std::string str_n = n.getNumber();

  for (char digit : str_n) {
    sum += digit - '0';
  }

  return sum;
}

int main () {

  unsigned largest = 0;

  for (int a=1; a < 100; a++) {
    for (int b=1; b < 100; b++) {
      unsigned sum = digit_sum(pow(a, b));
      if (sum > largest)
        largest = sum;
    }
  }

  std::cout << largest << "\n";

  return 0;
}
