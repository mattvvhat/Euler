import math

def factors (n):

  d = 2
  res = dict()
  while (d <= n):
    if (n % d == 0):
      n /= d
      if d in res:
        res[d] += 1
      else:
        res[d] = 1
    else:
      d += 1

  return res


print factors(28)
