#include <iostream>
#include <string>
#include <sstream>

#include <set>

#include "Eratosthenes.hpp"
#include "aljebr.h"

typedef int number;

using namespace std;

/**
 * Get Number of Permutations
 *
 */
set<number> prime_permutations (number n) {
  set<number> perms;
  string str = to_string(n);
  do {
    number val = atoi(str.c_str());
    if (Eratosthenes::is_prime(val))
      perms.insert(val);
  } while(next_permutation(str.begin(), str.end()));

  return perms;
}

/**
 * Get Map of Differences
 *
 */
map<int, set<number> > differences (set<number> perms) {
  map<int, set<number> > vals;

  for (auto i=perms.begin(); i != perms.end(); ++i) {
    for (auto j=i; j != perms.end(); ++j) {
      int diff = abs(*j - *i);

      if (diff == 0)
        continue;

      vals[diff].insert(*i);
      vals[diff].insert(*j);
    }
  }

  return vals;
}

/**
 * Filter Prime Number from Set
 */
void filter_primes (set<number> &vals) {
  for (auto it=vals.begin(); it != vals.end(); ++it) {
    if (!Eratosthenes::is_prime(*it))
      vals.erase(it);
  }
}

//
void print_set (const set<number> &vals) {
  for (auto it=vals.begin(); it != vals.end(); ++it) {
    cout << *it << " ";
  }
  cout << "\n";
}

// Go.

int main () {

  Eratosthenes::expand(10000);

  // ...
  for (auto it=Eratosthenes::begin(); it != Eratosthenes::end(); ++it) {
    number val = *it;

    if (val < 1000)
      continue;
    else if (val > 9999)
      break;

    auto perms = prime_permutations(val);
    auto diffs = differences(perms);

    for (auto i=diffs.begin(); i != diffs.end(); ++i) {

      unsigned k = i->first;
      set<number> list = i->second;

      for (auto j=list.begin(); j != list.end(); ++j) {
        if (list.count(*j+k) && list.count(*j+2*k)) {
          std::cout << *j << " " << *j+k << " " << *j+2*k << "\n";
        }
      }

    }

  }

  // End
  return 0;
}
