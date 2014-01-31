#include <iostream>
#include "aljebr.h"

#define TRI(n) (n*(n+1)/2)

// Typedefs... because this is a pain
typedef unsigned long ul;
typedef std::map<ul,ul> nmap;
typedef nmap::iterator nmap_iter;

int main () {

  ul largest_div_count = 0;
  ul i = 1;
  ul val;


  while (largest_div_count < 500) {
    int count = al::div_cnt(TRI(i));

    if (count > largest_div_count) {
      largest_div_count = count;
      val = TRI(i);
    }

    i++;
  }

  std::cout << "\n" << val << " " << largest_div_count << "\n";

  return 0;
}

