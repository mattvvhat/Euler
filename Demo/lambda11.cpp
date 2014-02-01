#include <iostream>
#include <functional>

int main () {
  auto func = [] () { std::cout << "yis\n"; };
  auto vals = [] (int x) -> void { std::cout << "(" << x << ")\n"; };
  auto rets = [] (int x, int y) -> int { return x*y; };
  std::function<int(int, int)> f = rets;
  std::cout << f(13, 2) << "\n";
  return 0;
}
