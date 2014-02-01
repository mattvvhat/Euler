def permutations_of( seq ):
  yield seq
  seqlen = len(seq)
  while True:
    j = seqlen - 2
    while j >= 0 and seq[j] >= seq[j+1]:
      j -= 1

    if j < 0: break

    i= seqlen - 1
    while i > j and seq[i] < seq[j]:
      i -= 1

    seq[j],seq[i] = seq[i],seq[j]
    seq[j+1:] = seq[-1:j:-1]
    yield seq

y = [ 1, 2, 3, 4 ]

k=0
#for x in permutations_of(y):
 # k+=1
  # print x

def rev (x, j):
  y = x
  return y[-1:j-1:-1]

print rev(y, 1)
