#include <iostream>
#include <set>
#include <algorithm>

#include "aljebr.h"

// ...
std::set<int> cycle (int n) {
  std::set<int> collection;
  char rawstr[256];
  sprintf(rawstr, "%d", n);
  std::string str = rawstr, orig(str);

  do {
    collection.insert(atoi(str.c_str()));
    std::rotate(str.rbegin(), str.rbegin()+1, str.rend());
  } while (str != orig);


  return collection;
}

// ...
int main () {
  std::set<int> primes;
  std::set<int> cyclic_primes;

  for (int i=2; i < 1000000; i++) {
    std::set<int> coll = cycle(i);
    bool cyclic_prime = true;
    for (auto it=coll.begin(); it != coll.end(); ++it) {
      if (cyclic_primes.count(*it) > 0) {
        continue;
      }
      else if (!al::is_prime(*it)) {
        cyclic_prime = false;
      }
    }

    if (cyclic_prime) {
      cyclic_primes.insert(i);
    }
  }

  int count=0;
  for (auto it=cyclic_primes.begin(); it != cyclic_primes.end(); ++it) {
    count++;
  }

  std::cout << count << "\n";

  // 

  return 0;
}
