#include <sstream>
#include <string>

namespace sugar {
  unsigned long long str_to_ull (std::string str) {
    unsigned long long n;
    std::istringstream (str) >> n;
    return n;
  }
}

namespace su = sugar;
