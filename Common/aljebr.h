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

  /**
   * Greatest Common Divisor
   * Compute the GCD of a natural number using Euclid's method.
   * @param a a natural number
   * @param b b natural number
   * @return the greatest common divisor of a and b
   */
  ul gcd (ul a, ul b) {
    if (b < a) {
      unsigned long int t = a;
      a = b;
      b = t;
    }

    unsigned long int r = b % a;
    unsigned long int m;

    while (r != 0) {
      b = a;
      a = r;
      m = b / a;
      r = b % a;
    }

    return a;
  }
  /**
   * Least Common Multiple
   * Compute the LCD of a natural number
   * @param a a natural number
   * @param b b natural number
   * @return the least common multiple of a and b
   */
  ul lcm (ul a, ul b) {
    return a*b/gcd(a, b);
  }
}

// Short name
namespace al = aljebr;
