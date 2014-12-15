#include <stdio.h>
#include <math.h>

#define DIVIDES(n,d) (n % d == 0)

// bool is_prime (n);

#define VAL 600851475143
// #define VAL 13195
// #define VAL 10

int main () {

  long unsigned int val = VAL;
  int d = 2;

  while (val > 1) {
    if (DIVIDES(val, d))
      val /= d;
    else
      d++;
  }

  printf("largest factor = %d\n", d);

  return 0;
}
