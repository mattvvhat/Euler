#include <iostream>
#include <cmath>
#include <map>

float angle (unsigned a, unsigned b, unsigned c) {
  float cosT;
  cosT = pow(a, 2) + pow(b, 2) - pow(c, 2);
  cosT = cosT / 2 / a / b;
  return cosT;
}

#define MAX 1000

int main () {

  std::map <unsigned, unsigned> counts;

  // ...
  for (int i=1; i <= MAX; i++) {
    for (int j=i; j <= MAX; j++) {
      for (int k=j; k <= MAX; k++) {
        float cosT = angle(i, j, k);

        if (cosT == 0) {
          unsigned p = i+j+k;
          counts[p] = counts.count(p) + 1;
        }
        else if (abs(cosT) > 1) {
          break;
        }
      }
    }
  }

  int largest=0, large_val=0;
  // ...
  for (auto it=counts.begin(); it != counts.end(); ++it) {
    if (it->second > largest) {
      largest = it->second;
    }
  }

  std::cout << largest << "\n";


  return 0;
}
