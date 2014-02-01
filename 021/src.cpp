#include <iostream>
#include "aljebr.h"

// Typedefs... because this is a pain
typedef unsigned long ul;
typedef std::map<ul,ul> nmap;
typedef nmap::iterator nmap_iter;

/**
 *
 */
void printvec (const std::vector <ul> & vec) {
  typedef std::vector<ul>::const_iterator const_iter;
  for (const_iter it=vec.begin(); it != vec.end(); ++it) {
    std::cout << *it << " ";
  }
  std::cout << std::endl;
}

/**
 * Whether is part of an amical pair
 */
bool is_amicable (unsigned long n) {
  ul other = al::divisor_sum(n);
  return n == al::divisor_sum(other) && n != other;
}

int main () {

  std::vector <ul> divisors;
  std::vector <ul> amicable;

  for (int i=0; i <= 10000; i++) {
    if (is_amicable(i)) {
      amicable.push_back(i);
    }
  }

  ul sum=0;
  typedef std::vector<ul>::const_iterator const_iter;
  for (const_iter it=amicable.begin(); it != amicable.end(); ++it) {
    sum += *it;
  }

  std::cout << sum << " ";

  return 0;
}

