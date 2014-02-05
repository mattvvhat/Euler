#include <iostream>

#include <string>

std::string to19[20] = {
  "", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
  "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"
};

std::string tens[10] = {
  "", "ten", "twenty", "thirty", "forty", "fifty",
  "sixty", "seventy", "eighty", "ninety"
};

std::string to_word (int n) {
  std::string str;

  if (n == 0)
    return "zero";

  else if (0 < n && n < 20)
    return to19[n];

  if (20 <= n && n < 100) {
    return tens[n / 10] + to19[n % 10];
  }

  if (100 <= n && n < 1000) {
    str = to19[n/100] + "hundred";

    if (n % 100 != 0)
      return str + "and" + to_word(n % 100);
    else
      return str;
  }
  
  if (n == 1000)
    return "onethousand";

  return str;
}

int main () {

  unsigned long sum=0;
  for (int i=1; i <= 1000; i++) {
    sum += to_word(i).length();
  }

  std::cout << sum << "\n";

  return 0;
}
