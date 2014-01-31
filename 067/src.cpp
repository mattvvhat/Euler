#include <fstream>
#include <iostream>
#include "aljebr.h"

#define TRI(n) (n*(n+1)/2)

// Typedefs... because this is a pain
typedef unsigned long ul;
typedef std::map<ul,ul> nmap;
typedef nmap::iterator nmap_iter;

using namespace std;

int main () {

  ifstream fin;
  fin.open("triangle.txt");
  unsigned long grid[100][100];

  // CLEAR BUFFER
  for (int i=0; i < 100; i++) {
    for (int j=0; j < 100; j++) {
      grid[i][j] = 0;
    }
  }

  int k=0;

  for (int i=0; i < 100; i++) {
    for (int j=0; j < i+1; j++) {
      if (fin.eof())
        break;

      unsigned long int n;
      fin >> n;
      grid[i][j] = n;
    }
  }

  for (int i=98; i >= 0; i--) {
    for (int j=0; j < 100-1; j++) {
      grid[i][j] += fmax(grid[i+1][j], grid[i+1][j+1]);
    }
  }

  unsigned long largest = 0;


  for (int j=0; j < 100; j++) {
    largest = grid[0][j] > largest ? grid[0][j] : largest;
  }

  cout << largest << "\n";


  return 0;
}

