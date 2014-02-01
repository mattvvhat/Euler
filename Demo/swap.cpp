#include <iostream>

template <typename type>
void fake_swap (type &x, type &y) {
  type t = x;
  x = y;
  y = t;
}

int main () {
  int a=3, b=4;
  fake_swap<int>(a, b);
  std::cout << a << ", " << b << "\n";
  return 0;
}
