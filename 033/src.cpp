#include <iostream>
#include <string>
#include "aljebr.h"

struct fract {
  int n;
  int d;
};

fract digit_reduct (int n, int d) {
  std::string nom = std::to_string(n);
  std::string den = std::to_string(d);

  for (auto i=nom.begin(); i != nom.end(); ++i) {
    for (auto j=den.begin(); j != den.end(); ++j) {
      if (*i == *j) {
        nom.erase(i, i+1);
        den.erase(j, j+1);
        return { atoi(nom.c_str()), atoi(den.c_str()) };
      }
    }
  }

  return { atoi(nom.c_str()), atoi(den.c_str()) };
}

fract reduct (int n, int d) {
  int divisor = al::gcd(n, d);
  return { (int)n/divisor, (int)d/divisor };
}

double asd (const fract & f) {
  double n = f.n, d = f.d;
  return n/d;
}

int main () {

  int mul=1;
  fract final;
  final.n = final.d = 1;

  for (int i=10; i < 100; i++) {
    for (int j=i+1; j < 100; j++) {
      fract left, right;
      left  = digit_reduct(i, j);
      right = reduct(i, j);

      double a = asd(left);
      double b = asd(right);

      if (a == b && left.n != i && left.d != j && right.n != j && right.d != j && i % 10 != 0 && j % 10 != 0) {
        final.n *= right.n;
        final.d *= right.d;
        // std::cout << a << " ?= " << b << "\n";
        std::cout << i << "/" << j << "\n";
        std::cout << left.n << "/" << left.d << "\n";
        std::cout << right.n << "/" << right.d << "\n";
        std::cout << "\n";
      }
    }
  }
  
  final = reduct(final.n, final.d);
  std::cout << "---\n" << final.n << "/" << final.d << "\n";
  std::cout << "answer = " << final.d << "\n";

  return 0;
}
