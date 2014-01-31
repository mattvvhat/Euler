#include <stdio.h>
#include <math.h>

#define VAL 1000

int main () {

  int y1=1, y2=2, y3=-1;

  int sum=0;

  for (int i=0; 1; i++) {
    y3 = y1+y2;
    y1 = y2;
    y2 = y3;

    if (y3 > 4000000) {
      break;
    }

    sum += y3;
  }

  printf("sum = %d\n", sum+2);

  return 0;
}
