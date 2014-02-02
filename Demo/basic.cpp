#include <iostream>
#include "Basic.hpp"

int main () {
  Basic<int> x(new int(7));
  Basic<int> y(new int(69));
  Basic<int> z(y);
  Basic<int> w(std::move(Basic<int>(NULL)));

  // Copy Assignment
  x = y;
  // Move Assignment
  y = Basic<int>(new int (7));
  // Copy Assignment
  y = x;

  return 0;
}
