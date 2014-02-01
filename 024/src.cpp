#include <iostream>
#include <algorithm>
#include <string>
#include <vector>

using namespace std;

// Defined later
void permute (vector<int> & vec);

void reverse_slice (vector<int> &, int, int);

int main () {

  std::string perm = "0123456789";

  int k=1;
  do {
    if (k == 1000000) {
      break;
    }
    k++;
  } while (next_permutation(perm.begin(), perm.end()));

  cout << perm << "\n";
  // permute(perm);

  return 0;
}

/**
 *
 */
void permute (vector<int> & seq) {
  int seqlen = seq.size();

  // INSIDE "LOOP"
  int j = seqlen - 2;

  copy(seq.begin(), seq.end(), ostream_iterator<int>(cout, " "));
  cout << "\n";

  while (true) {
    while (j >= 0 && seq[j] >= seq[j+1]) {
      j--;
    }

    if (j < 0)
      break;

    int i = seqlen - 1;
    while (i > j && seq[i] < seq[j]) {
      i--;
    }

    int t   = seq[j];
    seq[j]  = seq[i];
    seq[i]  = t;
    reverse_slice(seq, j+1, -1);

    copy(seq.begin(), seq.end(), ostream_iterator<int>(cout, " "));
    cout << "\n";
  }
}

/**
 *
 *
 */
void reverse_slice (vector<int> & vec, int a, int b) {
  if (b == -1) {
    b = vec.size();
  }

  if (a > b) {
    int t = a;
    a = b;
    b = t;
  }

  typedef vector<int>::iterator iter;
  iter left, right;
  left  = vec.begin() + a;
  right = vec.begin() + b;

  std::reverse(left, right);
}
