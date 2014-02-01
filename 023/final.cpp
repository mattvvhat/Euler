#include <iostream>
#include <fstream>
#include "aljebr.h"
#include <set>

#define MAX 28123

using namespace std;

int main () {
  ifstream in;
  in.open("abundant_numbers.txt");

  int n;


  set <int> abundant_sums;
  vector<int> abundants;

  while (!in.eof()) {
    in >> n;
    abundants.push_back(n);
  }

  typedef vector<int>::iterator const_iter;

  for (const_iter i=abundants.begin(); i != abundants.end(); ++i) {
    for (const_iter j=abundants.begin(); j != abundants.end(); ++j) {
      abundant_sums.insert(*i + *j);
    }
  }

  int summation = 0;
  for (int i =0; i <= MAX; i++) {
    if (abundant_sums.count(i) == 0) {
      summation += i;
    }
  }

  cout << summation << "\n";

  in.close();

  return 0;
}


