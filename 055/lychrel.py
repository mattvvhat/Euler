def palindrome (n):
  m=0
  while n > 0:
    m *= 10
    m += n % 10
    n /= 10
  return m

def is_lychrel (n):
  count=0
  while count < 50:
    count += 1
    n += palindrome(n)
    if n == palindrome(n):
      return False
  return True

count=0

for n in range(10000):
  if is_lychrel(n): count += 1


print count
