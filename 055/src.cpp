#include <iostream>
#include <string>

// /*
unsigned palindrome (unsigned n) {

  unsigned m=0;
  while (n > 0) {
    m *= 10;
    m += n % 10;
    n /= 10;
  }
  return m;
}
//*/

bool is_lychrel (unsigned long long n) {
  unsigned count=0;

  while (count++ < 50) {
    n += palindrome(n);

    if (n == palindrome(n)) {
      return false;
    }
  }

  return true;
}

int main () {

  std::cout << is_lychrel(10677) << "\n";
  unsigned count=0;
  for (unsigned i=10; i < 10000; i++) {
    if (is_lychrel(i))
      count++;
  }

  std::cout << count << "\n";

  return 0;
}
