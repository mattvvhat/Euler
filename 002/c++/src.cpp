#include <iostream>

int main () {

  unsigned y0=1, y1=2, y2;

  int sum = 0;

  do {
    y2 = y1+y0;
    y0 = y1;
    y1 = y2;

    if (y2 % 2 == 0)
      sum += y2;
  } while (y2 < 4000000);

  std::cout << sum+2 << "\n";


  return 0;
}
