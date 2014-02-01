#include <iostream>
#include <functional>

struct f {
  int operator()(int a, int b) {
    return a*a;
  };
};

int main () {
  f x;
  std::function<int(int, int)> f = x;
  std::cout << f(13, 2) << "\n";
  return 0;
}
