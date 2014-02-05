#include <iostream>
#include <string>
#include <set>

std::string build_mul (unsigned n) {
  std::string str;

  unsigned k=1;
  while (str.length() < 9) {
    str += std::to_string(k*n);
    k++;
  }

  return str;
}

bool is_pandigital (const std::string & str) {
  std::set<char> digits;
  for (auto c : str) {
    if (c == '0')
      return false;
    digits.insert(c);
  }
  return digits.size() == 9 && str.size() == 9;
}

int main () {

  unsigned largest = 0;
  unsigned lval = 0;

  for (int i=0; i < 9999; i++) {
    std::string mul = build_mul(i);
    unsigned val    = atoi(mul.c_str());

    if (is_pandigital(mul) && val > largest) {
      largest = val;
      lval = i;
    }
  }

  std::cout << largest << " (" << lval << ")\n";

  return 0;
}
