#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <cmath>

using namespace std;

int score (string str) {
  int sum=0;
  typedef string::const_iterator iter;
  for (iter it=str.begin(); it != str.end(); ++it) {
    sum += *it - 'A' + 1;
  }
  return sum;
}

bool is_tri (unsigned y) {
  float rad = sqrt(1 + 8*y);
  return rad == (int) rad;
}

int main () {

  vector<string> entries;

  ifstream ifs;
  ifs.open("words.txt");

  string entry="";
  while (!ifs.eof()) {
    char c[2] = { ifs.get(), '\0' };

    if (strcmp(c, "\"") == 0) {
      // no-op
    }
    else if (strcmp(c, ",") == 0) {
      entries.push_back(entry.c_str());
      entry = "";
    }
    else if (ifs.eof()) {
      entries.push_back(entry.c_str());
    }
    else {
      entry += c;
    }
  }

  ifs.close();

  int sum=0;
  typedef std::vector<std::string>::iterator iter;
  for (iter it=entries.begin(); it != entries.end(); ++it) {
    if (is_tri(score(*it)))
      sum++;
    // std::cout << *it << ",";
  }

  cout << sum << "\n";


  return 0;
}
