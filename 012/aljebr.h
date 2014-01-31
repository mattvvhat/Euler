#pragma once

#include <math.h>
#include <vector>
#include <map>

#define DIVIDES(n,d) (n % d == 0)

namespace aljebr {
  typedef unsigned long ul;
  typedef std::map<ul,ul> nmap;
  typedef nmap::iterator nmap_iter;

  /**
   * Whether a Number is Prime
   * Naive algorithm for determining whether a number is prime. *Many*
   * optimizations to this exist.
   * @param n a number to determine whether it is prime
   * @return  whether it is prime
   */
  bool is_prime (unsigned long int n) {
    unsigned long int d = 2;
    unsigned long int rootval = sqrt(n);

    while (d <= rootval) {
      if (n % d == 0)
        return false;
      d++;
    }

    return true;
  }

  /**
   * Prime Factorization
   * Return map
   * @param n the number to find the divisors of
   * @return p -> k, where p is a prime and k is the number of times divisible
   */
  std::map <unsigned long, unsigned long> factors (unsigned long n) {
    std::map <unsigned long, unsigned long> divs;
    unsigned long d = 2;

    while (d <= n) {
      if (n % d == 0) {
        n /= d;
        if (divs.count(d) == 0)
          divs[d] = 1;
        else
          divs[d]++;
      }
      else {
        d++;
      }
    }

    return divs;
  }

  /**
   * Divisor Function
   * Counts the number of divisors
   */
  int div_cnt (int n) {
    nmap factors = aljebr::factors(n);
    int m = 1;
    for (nmap_iter i=factors.begin(); i != factors.end(); ++i) {
      m *= i->second + 1;
    }
    return m;
  }
}

// Short name
namespace al = aljebr;
