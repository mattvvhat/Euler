#include <iostream>
#include <map>
#include <cmath>

#define IS_INTEGER(n) (n == (unsigned long long) n)

typedef unsigned long long number;

int main () {

  std::map<number, number> perims;

  for (unsigned a=1; a <= 500; a++) {
    for (unsigned b=a; b <= 500; b++) {
      double c = sqrt(pow(a, 2) + pow(b, 2));
      if (c == 1) {
        std::cout << "!?";
      }
      if (IS_INTEGER(c)) {
        unsigned p = a+b+c;
        if (perims.count(p))
          perims[p] = perims[p] + 1;
        else
          perims[p] = 1;
      }
    }
  }

  unsigned largest=0;
  unsigned val=0;

  for (auto it=perims.begin(); it != perims.end(); ++it) {
    if (it->second > largest && it->first <= 1000) {
      largest = it->second;
      val = it->first;
    }
  }

  std::cout << largest << " " << val << "\n";

  return 0;
}
