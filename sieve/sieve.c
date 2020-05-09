#define _GNU_SOURCE

#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

// #include "what.h"


#define PRIME_COUNT 10000000


int main() {


  // Kinda like a map...
  bool *sieve = malloc(PRIME_COUNT * sizeof(bool));

  // Initiate all as potential primes
  for (int64_t i=0; i < PRIME_COUNT; i++) {
    sieve[i] = true;
  }

  sieve[0] = false;
  sieve[1] = false;

  // ...
  for (int64_t i=2; i < PRIME_COUNT; i++) {

    if (!sieve[i]) {
      continue;
    }

    int64_t j = i*i;

    while (j < PRIME_COUNT) {
      sieve[j] = false;
      j += i;
    }
  }


  int count = 0;
  for (int i=0; i < PRIME_COUNT; i++) {
    if (sieve[i]) {
      // printf("%d ", i);
      count++;
    }
  }



  printf("COUNT => %d", count);
  free(sieve);




  // You  know what to do.
  printf("\n\n");
  printf("(>x_x)> Fuck it\n");
  exit(EXIT_SUCCESS);
}
