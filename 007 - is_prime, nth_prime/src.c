#include <stdio.h>
#include <math.h>

#define VAL 10001

int is_prime (unsigned long int n);
unsigned long int nth_prime (int n_prime);

int main () {

  printf("%d -> %lu\n", VAL, nth_prime(VAL));

  return 0;
}

/**
 *
 *
 *
 */
int is_prime (unsigned long int n) {
  unsigned long int d = 2;
  unsigned long int rootval = sqrt(n);

  while (d <= rootval) {
    if (n % d == 0)
      return 0;
    d++;
  }

  return 1;
}

/**
 *
 *
 */
unsigned long int nth_prime (int n_prime) {
  int i = 2;
  int prime_count = 0;
  unsigned long int last_prime;

  while (prime_count != n_prime) {
    if (is_prime(i)) {
      prime_count++;
      last_prime = i;
    }
    i++;
  }

  return last_prime;
}
