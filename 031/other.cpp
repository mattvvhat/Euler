#include <iostream>

int main () {
  int target = 200;
  int coinSizes[8] = { 1, 2, 5, 10, 20, 50, 100, 200 };
  int ways[201];
  for (int i=0; i <= 200; i++) {
    ways[i] = 0;
  }
  ways[0] = 1;
   
  for (int i = 0; i < 8; i++) {
    for (int j = coinSizes[i]; j <= target; j++) {
      ways[j] += ways[j - coinSizes[i]];
    }
  }

  for (int i=0; i <= 200; i++) {
    // std::cout << i << " -> " << ways[i] << "\n";
  }

  std::cout << ways[9] << "\n";
  std::cout << ways[4] << " ";
  std::cout << ways[3] << " ";
  std::cout << "\n";

  return 0;
}
