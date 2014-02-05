#include <iostream>
#include <cmath>
#include <string>

bool is_integer (double n) {
  return (int) n == n;
}

int main () {

  unsigned i=10, n=10;
  std::string val;

  unsigned long m=1;
  while (i <= 1000000) {

    val = std::to_string(n);

    for (char c : val) {
      // std::cout << c << " (" << i << ")\n";
      if (is_integer(log10(i)))
        m *= c - '0';
      i++;
    }

    n++;
  }

  std::cout << m << "\n";

  return 0;
}
