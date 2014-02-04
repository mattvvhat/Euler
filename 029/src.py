
what = {}

for a in range(2, 101):
  for b in range(2, 101):
    what[a**b] = 1


print len(what)
