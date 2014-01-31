#include <stdio.h>
#include <math.h>
#include <string.h>

int main () {

  for (int a=1; a < 999; a++) {
    for (int b=01; b < 999; b++) {
      int c2 = pow(a, 2) + pow(b, 2);
      float c = sqrt(c2);

      if (a+b+c==1000) {
        printf("%f\n", a*b*c);
        break;
      }
    }
  }

  return 0;
}

