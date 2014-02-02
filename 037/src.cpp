#include <iostream>
#include <set>
#include "Eratosthenes.hpp"

using namespace std;

set<int> truncations (int n) {
  set<int> truncs;
  string val;

  val = to_string(n);
  while (val.length()) {
    truncs.insert(atoi(val.c_str()));
    val.erase(0, 1);
  }

  val = to_string(n);
  while (val.length()) {
    truncs.insert(atoi(val.c_str()));
    val.erase(val.length()-1, 1);
  }
  return truncs;
}

bool is_cyclic_prime (int n) {
  set<int>truncs = truncations(n);
  for (set<int>::iterator it=truncs.begin(); it != truncs.end(); ++it) {
    if (!Eratosthenes::is_prime(*it))
      return false;
  }
  return true;
}

int main () {

  cout << "* creating sieve\n";
  Eratosthenes::expand(2000000);

  cout << "* sieve created\n";

  string str = to_string(300);
  int n = atoi(str.c_str());

  int sum=0, count=0;;
  for (int i=11; i < 2000000; i++) {
    if (is_cyclic_prime(i)) {
      sum += i;
      count++;
    }
  }

  cout << "sum   = " << sum << "\n";
  cout << "count = " << count << "\n";



  return 0;
}
