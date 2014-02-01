#include <iostream>
#include <algorithm>
#include <map>

// Problem 31 - Dynamic programming

int main () {

  typedef unsigned long long ull;
  std::map<ull, ull> combinations;

  // Initial known number of combinations
  // for (int i=0; i < 200; i++)
  //   combinations[i] = 0;
  combinations[  0] = 1;
  combinations[  1] = 1;
  combinations[  2] = 1;
  combinations[  5] = 1;
  combinations[ 10] = 1;
  combinations[ 20] = 1;
  combinations[ 50] = 1;
  combinations[100] = 1;
  combinations[200] = 1;

  // 
  for (int i=1; i <= 200; i++) {
    if (combinations.count(i) == 0)
      combinations[i] = 0;

    for (int j=1; j <= i/2; j++) {
      combinations[i] += combinations[i-j];
    }
  }

  // Output
  for (auto it = combinations.begin(); it != combinations.end(); ++it) {
    // std::cout << it->first << " -> " << it->second << "\n";
  }

  std::cout << combinations[200] << "\n";

  // 73682
  return 0;
}
