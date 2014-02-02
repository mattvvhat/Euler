import math

def inv_tri (y):
  np = +(-1 + math.sqrt(1 + 8*y))/2.
  nm = +(-1 - math.sqrt(1 + 8*y))/2.
  return np, nm

def is_tri (y):
  rad = math.sqrt(1+8*y)
  return rad == math.floor(rad)

print inv_tri(1)
print inv_tri(3)
print inv_tri(10)
print inv_tri(55)
print inv_tri(56)
print
print is_tri(1)
print is_tri(3)
print is_tri(10)
print is_tri(55)
print is_tri(56)
