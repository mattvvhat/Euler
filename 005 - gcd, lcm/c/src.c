#include <stdio.h>
#include <math.h>

#define DIVIDES(n,d) (n % d == 0)

unsigned long int gcd (unsigned long int a, unsigned long int b);
unsigned long int lcm (unsigned long int a, unsigned long int b);

int main () {

  unsigned long int val = 1;

  for (int i=2; i < 20; i++) {
    val = lcm(val, i);
  }

  printf("%lu\n", val);
  return 0;
}

unsigned long int gcd (unsigned long int a, unsigned long int b) {
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

unsigned long int lcm (unsigned long int a, unsigned long int b) {
  return a*b/gcd(a, b);
}
