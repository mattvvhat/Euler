#include <iostream>
#include <set>
#include <map>
#include <cmath>
#include <vector>

#define LEN(n) (n > 0 ? (int) log10(n)+1 : 1)


bool is_panmul (unsigned a, unsigned b) {

  std::string p = std::to_string(a);
  std::string q = std::to_string(b);
  std::string r = std::to_string(a*b);

  std::set<unsigned> digits;

  for (auto ch : p+q+r) {
    if (ch == '0')
      return false;
    digits.insert(ch);
  }

  return digits.size() == 9 && (p.size() + q.size() + r.size() == 9);
}


int main () {

  std::set<unsigned> pans;

  unsigned i=0, j=0;
  unsigned long res;

  unsigned int sum=0;

  for (unsigned a=1; a < 999; a++) {
    for (unsigned b=a+1; LEN(a) + LEN(b) + LEN(a*b) < 10; b++) {
      if (is_panmul(a, b)) {
        std::cout << a << " x " << b << " = " << a*b << "\n";
        pans.insert(a*b);
      }

    }
  }

  for (auto it=pans.begin(); it != pans.end(); ++it) {
    sum += *it;
  }

  std::cout << sum << "\n";
  std::cout << "\nFine\n";

  return 0;
}
