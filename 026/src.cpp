#include <iostream>
#include <string>
#include <cmath>
#include <map>

/**
 * SOLUTION: develop division for strings
 */

int cycle_count (unsigned d) {
  std::map<std::string, unsigned> memory;
  std::string result;
  int power=1, val=1, count=0, place=0;

  while (true) {

    // Custom container to remember
    std::string operation;

    unsigned k=1, r, shifts=0;

    while (d > val) {
      val *= 10;
      shifts++;
    }

    unsigned digit = val/d;
    r   = val % d;
    val = val - digit * d;

    // INCREMEMTED PLACE AND ADD TO MEMORY
    operation = std::to_string(val) + " / " + std::to_string(d);
    if (memory.count(operation)) {
      return place-memory[operation];
      // return result + "~";
    }
    memory[operation] = place++;

    // Add to string
    while (--shifts)
      result += '0';

    result += std::to_string(digit);


    // CATCH END CONDITIONS
    if (r == 0 || count++ > 400000)
      break;
  }

  return 0;
  // return result;
}

int main () {

  // std::cout << cycle_count(2) << "\n";
  // std::cout << cycle_count(4) << "\n";
  // std::cout << cycle_count(8) << "\n";
  // std::cout << cycle_count(80) << "\n";
  // std::cout << cycle_count(3) << "\n";

  int largest=0;
  int val=0;

  for (int i=2; i <= 1000; i++) {
    int count = cycle_count(i);
    if (count > largest) {
      largest = count;
      val = i;
    }
  }

  std::cout << val << " (" << largest << ")\n";

  return 0;
}
