#include <stdio.h>
#include <math.h>

#define VAL 1000

int divisor_sum (int n);

int main () {

  int result = divisor_sum(VAL);

  printf("%d", result);

  return 0;
}

int divisor_sum (int n) {
  int sum = 0;

  for (int i=0; i < n; i++) {
    if (i % 3 == 0 || i % 5 == 0) {
      sum += i;
    }
  }

  return sum;
}
