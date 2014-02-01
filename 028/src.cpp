#include <iostream>

using namespace std;

#define GRID_SIZE 1001

int main () {

  auto ur = [&] (int n) -> int { return (2*n+1)*(2*n+1); };
  auto ul = [&] (int n) -> int { return ur(n) - 2*n; };
  auto bl = [&] (int n) -> int { return ul(n) - 2*n; };
  auto br = [&] (int n) -> int { return bl(n) - 2*n; };

  // init to 1, because we're not counting the first term ourselves
  unsigned sum = 1;

  for (int i=1; i < (GRID_SIZE-1)/2+1; i++) {
    sum += ur(i) + ul(i) + bl(i) + br(i);
  }

  cout << sum << "\n";

  return 0;
}
