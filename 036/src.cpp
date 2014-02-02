#include <iostream>
#include <string>

bool is_pal2 (int n) {
  std::string str = "";

  while (n != 0) {
    char rawstr[255];
    sprintf(rawstr, "%d", n%2);
    str += rawstr;
    n /= 2;
  }

  for (int i=0, len=str.length(); i < len/2; i++) {
    if (str[i] != str[len-i-1])
      return false;
  }

  return true;
  return true;
}

bool is_pal10 (int n) {
  char rawstr[25];
  sprintf(rawstr, "%d", n);
  std::string str = rawstr;

  for (int i=0, len=str.length(); i < len/2; i++) {
    if (str[i] != str[len-i-1])
      return false;
  }

  return true;
}

int main () {

  int sum=0;
  for (int i=1; i < 1000000; i++) {
    if (is_pal2(i) && is_pal10(i)) {
      sum += i;
    }
  }

  std::cout << sum << "\n";


  return 0;
}
