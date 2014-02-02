#include <iostream>
#include <string>
#include <algorithm>
#include <sstream>

int main () {
  typedef std::string::iterator string_iter;

  std::string pan = "0123456789";
  int primes[7] = { 2, 3, 5, 7, 11, 13, 17 };

  unsigned long long sum = 0;

  do {
    bool has_property = true;

    for (int i=1; i < pan.length()-2; i++) {
      unsigned long n = atoi(pan.substr(i, 3).c_str());
      unsigned long d = primes[i-1];

      if (n % d != 0) {
        has_property = false;
        break;
      }

    } 

    if (has_property) {
      unsigned long n;
      std::stringstream(pan) >> n;
      sum += n;
    }

  } while (std::next_permutation(pan.begin(), pan.end()));

  std::cout << "sum = " << sum << "\n";

  return 0;
}
