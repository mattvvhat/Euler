#!/usr/bin/python

def digit_sum (n):
  s=0

  while n > 0:
    s += n % 10
    n /= 10

  return s

print digit_sum(2**1000)
