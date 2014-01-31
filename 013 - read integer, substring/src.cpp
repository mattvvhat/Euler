#include <iostream>
#include <fstream>
#include "aljebr.h"
#include "sugar.h"

// Typedefs... because this is a pain
typedef unsigned long ul;
typedef unsigned long long ull;
typedef std::map<ul,ul> nmap;
typedef nmap::iterator nmap_iter;

int main () {

  std::ifstream ifs;
  ifs.open("numbers.txt");

  ull val=0;
  int line_count = 0;

  while (!ifs.eof()) {
    if (line_count == 100)
      break;
    line_count++;

    // Get entire line
    char rawline[255];
    ifs.getline(rawline, 255);

    // Substring of rawline
    std::string line = rawline;
    line = line.substr(0, 15);

    // convert to ull
    ull n = su::str_to_ull(line);

    // NOTE: n is now equivalent to (int) rawline / 10^40

    std::cout << n << "\n";
    val += n;
    // 

  }

  std::cout << val << "\n";

  ifs.close();
  return 0;
}

