#include <iostream>
#include <algorithm>
#include <vector>
#include <string>

#include "Eratosthenes.hpp"
#include "aljebr.h"

typedef unsigned number;
typedef std::vector<unsigned> number_list;

/**
 * Combinations of Size K from the Iteratable Object
 *
 */
std::vector<number_list> combinations (unsigned k, const std::vector<number> &vec) {
  // Final list-of-lists to return
  std::vector<number_list> combos;
  // Which values to transfer over
  std::vector<bool> yesno(vec.size(), false);

  // Initial block of ones
  for (unsigned i=0; i < k; i++)
    yesno[i] = true;

  // Permute yesno
  do {
    // Create a new copy to push to the list
    std::vector<number> copy;
    // Copy over each relevant value
    for (int k=0; k < yesno.size(); k++) {
      if (yesno[k] == true)
        copy.push_back(vec[k]);
    }
    // Push copy to the return value
    combos.push_back(std::move(copy));
  } while (std::prev_permutation(yesno.begin(), yesno.end()));

  return combos;
}

std::vector<number> range (unsigned n) {
  std::vector<number> vals(n, 0);
  for (unsigned i=0; i < n; i++)
    vals[i] = i;
  return vals;
}

#define SIZE 99999

using namespace std;

unsigned count_primes (unsigned n) {

  string str_n = to_string(n);
  number_list positions = range(str_n.size());

  unsigned largest = 0;

  for (unsigned k=1; k < str_n.size(); k++) {

    vector<number_list> combos = combinations(k, range(str_n.size()));

    // Enumerate possible combinations.
    // This can actually just be assumed to be 8... but whatever
    for (auto el : combos) {

      string str_stars(str_n);

      for (auto pos : el)
        str_stars[pos] = '*';

      unsigned count=0;
      for (unsigned digit=0; digit <= 9; digit++) {
        string str_copy(str_stars);
        for (auto it=str_copy.begin(); it != str_copy.end(); it++) {
          if (*it == '*')
            *it = digit + '0';
        }
        if (al::is_prime(atoi(str_copy.c_str())))
          count++;
      }

      if (count == 8) {
        cout << str_stars << "\n";
      }

      if (count > largest)
        largest = count;
      
      // cout << str_stars << " (" << count << ")\n";

    }

  }


  return largest;
}

int main () {

  Eratosthenes::expand(300000);

  unsigned largest = 0;
  unsigned val;

  // cout << count_primes(100109) << "\n";
  // return 0;

  for (unsigned prime : Eratosthenes::copy()) {
    string str_n = to_string(prime);
    for (unsigned k=1; k < str_n.size(); k++) {

      vector<number_list> list = combinations(k, range(str_n.size()));

      for (auto combo : list) {
        // Star-out the values
        string str_stars(str_n);
        for (unsigned p : combo) {
          str_stars[p] = '*';
        }

        unsigned count=0;
        { // Replace Digits
          for (unsigned digit=0; digit <= 9; digit++) {
            if (combo[0] == 0 && digit == 0)
              continue;

            string str_copy(str_stars);
            for (unsigned p=0; p < str_copy.size(); p++)
              if (str_copy[p] == '*')
                str_copy[p] = digit + '0';

            if (al::is_prime(atoi(str_copy.c_str())))
              count++;
          }
        }
        if (count > largest) {
          largest = count;
          val = prime;
        }
      }

    }
  }

  std::cout << val << " (" << largest << ")\n";

  return 0;
}
