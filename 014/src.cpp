#include <iostream>

// Typedefs... because this is a pain
typedef unsigned long ul;

// Collatz
ul collatz (ul n) {
  ul count=0;
  while (n != 1) {
    n = n % 2 ? 3*n+1 : n/2;
    count++;
  }
  return count;
}

int main () {

  collatz(13);

  ul largest = 0;
  int lrg_val = 0;

  for (int i=1; i < 1000000; i++) {
    ul val = collatz(i);
    if (val > largest) {
      largest = val;
      lrg_val = i;
    }
  }

  std::cout << lrg_val << "\n";

  return 0;
}

