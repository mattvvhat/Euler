#include <stdio.h>
#include <math.h>

#define VAL 100

int main () {

  unsigned int long sum_of_sqrs = 0;
  unsigned int long sqr_of_sums = 0;

  for (int i=1; i <= VAL; i++) {
    sum_of_sqrs += i*i;
  }

  sqr_of_sums = VAL * (VAL+1) / 2;
  sqr_of_sums *= sqr_of_sums;

  

  printf("%lu\n", sqr_of_sums - sum_of_sqrs);

  return 0;
}
