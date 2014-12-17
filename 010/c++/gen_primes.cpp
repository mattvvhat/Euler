#include <iostream>
#include <fstream>
#include <math.h>

#include "aljebr.h"

using namespace std;

int main () {

  std::ofstream ofs ("primes.txt", std::ofstream::out);

  int i = 2;
  int prime_count = 0;
  unsigned long int sum = 0;
  unsigned long int n_prime = 0;

  while (i < 2000000) {
    if (al::is_prime(i)) {
      n_prime++;
      sum += i;

      ofs << i << "\n";
    }
    i++;
  }

  cout << "sum = " << sum << "\n";

  return 0;
}
