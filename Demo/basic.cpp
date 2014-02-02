#include <iostream>
#include <fstream>
#include "Basic.hpp"

int main () {
  Basic <int> x(new int(7)), y(new int(69)), z(y);
  Basic <int> w(std::move(Basic<int>(NULL)));

  x = y;
  y = Basic<int>(new int (7));
  y = x;

  std::cout << y << "\n";

  return 0;
}
