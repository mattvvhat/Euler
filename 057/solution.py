# Greatest Common Divisor
# Computes the Greatest Common Divisor between two integers using Euclid's
# Algorithm. If the innards are wrapped in more efficient methods, then it'd
# likely work quickly for very large integers!
def gcd (a, b):
  # Let us compute the gcd of negative numbers
  if a < 0: a = -a
  if b < 0: b = -b
  # Ensure that b is the larger number
  if b < a: (a, b) = (b, a)

  # Do the first iteration by hand
  r = b % a

  while r != 0:
    # Note: b = a * m + r
    b = a
    a = r
    m = b // a
    r = b % a

  return a

# Least Common Multiple
# Computes the least common multiple between two integers. This method is not
# good for very large integers
def lcm (a, b): return (a*b)/gcd(a, b)

def add (lhs, rhs):
  result = fract()
  scale = lcm(lhs.d, rhs.d)
  result.n = lhs.n * scale/lhs.d + rhs.n * scale/rhs.d
  result.d = scale
  return result.reduced()

class fract:
  def __init__ (self, n=1, d=1):
    self.n = n
    self.d = d
  def __repr__ (self):
    return "%d/%d" % (self.n, self.d)
  def reduced (self):
    div = gcd(self.n, self.d)
    return fract(self.n/div, self.d/div)
  def inverted (self):
    return fract(self.d, self.n)

# Start
summa = fract(2, 1)
two   = fract(2, 1)
half  = fract(1, 2)
one   = fract(1, 1)

count = 0

for i in range(1000):
  summa = add(two, summa.inverted())
  val = add(one, summa.inverted())
  val = val.reduced()
  nom = str(val.n)
  den = str(val.d)
  if len(nom) > len(den):
    count += 1

print count
