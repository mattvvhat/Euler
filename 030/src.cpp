#include <iostream>
#include <cmath>
#include <string>

unsigned digit_power_sum (unsigned n) {
  unsigned sum=0;
  std::string str = std::to_string(n);


  for (auto c : str) {
    unsigned n = c - '0';
    sum += pow(n, 5);
  }

  return sum;
}

int main () {

  unsigned sum=0;
  for (int i=2; i < 1000000; i++) {
    if (i == digit_power_sum(i)) {
      sum += i;
    }
  }

  std::cout << sum << "\n";

  return 0;
}
