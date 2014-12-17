namespace aljebr {
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
}

namespace al = aljebr;
