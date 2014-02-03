#include <iostream>
#include <cmath>

#define PEN(n) n*(3*n-1)/2

typedef unsigned long number;

bool is_pen (number y) {
  double n = sqrt(1 + 24*y);
  return (n == floor(n)) && ((number) (1+n) % 6 == 0);
}

int main () {

  for (int i=1; true; i++) {
    int pen_i = PEN(i);

    for (int j=1; j < i+1; j++) {
      int pen_j = PEN(j);

      int sum = pen_i + pen_j;
      int dif = pen_i - pen_j;

      if (is_pen(sum) && is_pen(dif)) {
        std::cout << dif << "\n";
        break;
      }

    }
  }

  return 0;
}
