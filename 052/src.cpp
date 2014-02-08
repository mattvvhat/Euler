#include <iostream>
#include <string>
#include <set>

bool is_permuted (unsigned m, unsigned n) {
  std::set<unsigned> digits_m, digits_n;
  std::string str_m = std::to_string(m), str_n = std::to_string(n);

  for (auto it=str_m.begin(); it != str_m.end(); ++it)
    digits_m.insert(*it - '0');

  for (auto it=str_n.begin(); it != str_n.end(); ++it)
    digits_n.insert(*it - '0');

  return digits_m == digits_n && str_m.length() == str_n.length();
}

int main () {

  unsigned x1=1;
  while (true) {
    unsigned x2 = 2 * x1;
    unsigned x3 = 3 * x1;
    unsigned x4 = 4 * x1;
    unsigned x5 = 5 * x1;
    unsigned x6 = 6 * x1;
    if (is_permuted(x1, x2) && is_permuted(x1, x2) && is_permuted(x1, x3) && is_permuted(x1, x4) && is_permuted(x1, x5) && is_permuted(x1, x5) && is_permuted(x1, x6)) {
      break;
    }
    x1++;
  }

  std::cout << x1 << "\n";

  return 0;
}
