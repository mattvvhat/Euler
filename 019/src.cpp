#include <iostream>

using namespace std;

bool leap(int y)
{
    return y % 400 ? y % 100 ? y % 4 ? false : true : false : true;
}

int monthdays(int m, int y)
{
    switch (m)
        {
              case 1: return 31;
                          case 2: return leap(y) ? 29 : 28;
                                      case 3: return 31;
                                                  case 4: return 30;
                                                              case 5: return 31;
                                                                          case 6: return 30;
                                                                                      case 7: return 31;
                                                                                                  case 8: return 31;
                                                                                                              case 9: return 30;
                                                                                                                          case 10: return 31;
                                                                                                                                       case 11: return 30;
                                                                                                                                                    case 12: return 31;
                                                                                                                                                                 default: return 0;
                                                                                                                                                                            }
}

int main()
{
    int d = 2, m = 1, y = 1901, ans = 0;
      
      while (y < 2001)
          {
                if (!d) ans++;
                    d = (monthdays(m, y) % 7 + d) % 7;
                        if (m == 12)
                              {
                                       m = 1;
                                              y += 1;
                                                  } 
                            else m += 1;
                              } 
        cout << ans << endl;
}
