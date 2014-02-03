#include <set>
#include <cmath>

class Eratosthenes {
public:

  Eratosthenes () {};

  static void expand (int n) {
    int rootn = sqrt(n);
    for (int i=2; i <= n; i++) {
      primes.insert(i);
    }
    max = n;

    for (int i=2; i <= rootn; i++) {
      int k=2;
      if (primes.count(i) == 0)
        continue;
      while (i*k <= n) {
        primes.erase(k*i);
        k++;
      }
    }
  }

  static bool is_prime (int n) {
    if (n > max) {
      expand(n);
    }
    return primes.count(n) != 0;
  }

  static std::set<int> copy () {
    std::set<int> s(primes);
    return s;
  }

  static std::set<int>::const_iterator begin () {
    return primes.begin();
  }

  static std::set<int>::const_iterator end () {
    return primes.end();
  }

private:
  static int max;
  static std::set<int> primes;
};

int Eratosthenes::max = 0;
std::set<int> Eratosthenes::primes = std::set<int>();
