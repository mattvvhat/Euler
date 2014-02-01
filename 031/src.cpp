#include <iostream>
#include <algorithm>
#include <vector>
#include <map>

// Problem 31 - Dynamic programming

int main () {

  typedef unsigned long long ull;
  std::map<ull, ull> values;
  std::vector<int> coins = { 1, 2, 5, 10, 20, 50, 100, 200 };

  // Initial known number of combinations
  for (int i=0; i <= 200; i++)
    values[i] = 0;

  values[1] = 1;

  // Solve all the number of times a coin goes into a value
  for (auto coinit = coins.begin(); coinit != coins.end(); ++coinit) {
    for (int i=*coinit; i <= 200; i++) {
      values[i] += values[i - *coinit];
    }
  }

  // Output
  std::cout << values[200] << "\n";

  // 73682
  return 0;
}
