#include <iostream>
#include <cmath>

typedef unsigned long long ull;

long long get_tri (ull y) {
  double root = sqrt(1+8*y);
  ull n = root;
  if (root != n)
    return -1;
  n = -1 + n;
  if (n % 2 != 0) {
    return -1;
  }
  return n/2;
}

bool is_tri (ull y) {
  double n = sqrt(1 + 8*y);
  return (n == floor(n)) && ((ull) (-1+n) % 2 == 0);
}

bool is_pen (ull y) {
  double n = sqrt(1 + 24*y);
  return (n == floor(n)) && ((ull) (1+n) % 6 == 0);
}

bool is_hex (ull y) {
  double n = sqrt(1 + 8*y);
  return (n == floor(n)) && ((ull) (1+n) % 4 == 0);
}

int main () {

  unsigned long long p = 40755;
  unsigned long long n = p+1;

  while (!(is_tri(n) && is_pen(n) && is_hex(n))) {
    if (n < p) {
      std::cout << "OVERFLOW!!!\n";
      break;
    }
    n++;
  }

  std::cout << "\n";
  std::cout << n << "\n";
  std::cout << n+1 << "\n";
  std::cout << get_tri(n) << "\n";
  std::cout << get_tri(n+1) << "\n";
  std::cout << is_tri(n) << "\n";
  std::cout << is_pen(n) << "\n";
  std::cout << is_hex(n) << "\n";
  // std::cout << n << "\n";

  return 0;
}
