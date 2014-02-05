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

    if (n < 2)
      return 0;

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
   * @param n the number to find the factors of
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
   * Divisors
   * Return list of divisors
   * @param n the number to find the divisors of
   * @return p -> k, where p is a prime and k is the number of times divisible
   */
  std::vector <ul> divisors (ul n, bool proper=true) {
    std::vector <ul> result;

    for (ul i=1; i < n; i++) {
      if (n % i == 0)
        result.push_back(i);
    }

    if (!proper)
      result.push_back(n);

    return result;
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
   * Divisor Sum
   * Compute the sum of the divisors of a number n
   * @param n a natural number to compute the divisor sum of
   * @param k a power to raise the divisor to
   */
  ul divisor_sum (unsigned long n, ul k=1) {
    typedef std::vector<ul>::const_iterator const_iter;
    std::vector <ul> divisors = aljebr::divisors(n);
    ul sum=0;
    for (const_iter it=divisors.begin(); it != divisors.end(); ++it) {
      sum += pow(*it, k);
    }
    return sum;
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

  /**
   * Factorial n!
   * Compute the factorial of an unsigned integer
   * @param n an unsigned integer
   * @return n!
   */
  ul factorial (ul n) {
    ul product = 1;
    while (n > 1) {
      product *= n--;
    }
    return product;
  }

  /**
   * Combinations (n, k)
   * Count the number of ways of selecting k from n
   * @param n the total population
   * @param k the number of things to select
   * @return k-combinations from n
   */
  ul combine (ul n, ul k) {
    ul product=1, stop=n-k;
    while (n > stop) {
      product *= n--;
    }
    while (k > 1) {
      product /= k--;
    }
    return product;
  }
}

// Short name
namespace al = aljebr;
