#include <iostream>
#include <cmath>

typedef unsigned long number;

number powmod(number n, number k, number mod) {
  number result = 1;
  while (k > 0) {
    result *= n % mod;
    result %= mod;
    k--;
  }
  return result;
}

int main () {
  std::cout << powmod(5, 5, 10) << "\n";
  unsigned long sum = 0;
  for (int i=1; i <= 1000; i++) {
    sum += powmod(i, i, pow(10, 11));
  }

  // Last 10 digits of this are the answer
  std::cout << "The last 10 digits of the following are the answer:\n";
  std::cout << sum << "\n";

  return 0;
}
