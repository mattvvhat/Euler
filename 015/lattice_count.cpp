#include <iostream>

typedef unsigned long ul;

#define SIZE 21

ul glob_grid[SIZE][SIZE];

int main () {

  for (int i=0; i < SIZE; i++) {
    for (int j=0; j < SIZE; j++) {
      if (i == 0 || j == 0)
        glob_grid[i][j] = 1;
      else
        glob_grid[i][j] = glob_grid[i-1][j] + glob_grid[i][j-1];
    }
  }

  std::cout << glob_grid[SIZE-1][SIZE-1] << "\n";

  return 0;
}
