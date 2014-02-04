#include <iostream>
#include <string>
#include <set>
#include "lib/BigInteger.h"

typedef unsigned long long int number;

// ...
BigInteger raised (unsigned a, unsigned b) {
  BigInteger val(1);
  while (b > 0) {
    val *= a;
    b--;
  }
  return val;
}

// ...
int main () {

  std::set<string> collection;

  for (int a=2; a <= 100; a++) {
    for (int b=2; b <= 100; b++) {
      collection.insert(raised(a, b).getNumber());
    }
  }

  std::cout << collection.size() << "\n";

  return 0;
}
