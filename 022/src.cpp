#include <iostream>
#include <fstream>
#include <vector>

int score (const std::string & str) {
  typedef std::string::const_iterator const_iter;
  int sum=0;
  for (const_iter it=str.begin(); it != str.end(); ++it) {
    char c = *it;
    sum += c - 'A' + 1;
  }
  return sum;
}

int main () {

  // Declare and open stream
  std::ifstream ifs;
  ifs.open("names.txt");

  std::string entry;
  std::vector <std::string> entries;

  char c[2];
  while (!ifs.eof()) {
    c[0] = ifs.get();
    c[1] = '\0';

    if (strcmp(c, "\"") == 0) {
      // no-op
    }
    else if (strcmp(c, ",") == 0) {
      entries.push_back(entry);
      entry = "";
    }
    else if (ifs.eof()) {
      entries.push_back(entry);
    }
    else {
      entry += c;
    }
  }

  // Pulling from a list...

  std::sort(entries.begin(), entries.end());

  std::cout << entries[938-1] << "\n";
  std::cout << score(entries[938-1]) << "\n";

  int grand_total = 0;
  for (int i=0; i < entries.size(); i++) {
    int total = score(entries[i]);
    total *= i+1;
    grand_total += total;
  }

  std::cout << grand_total << "\n";

  // Close stream
  ifs.close();

  return 0;
}

