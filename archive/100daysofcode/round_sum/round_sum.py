a = 16
b = 16
c = 16
num = (a, b, c)
def round_sum(a, b, c):

  print(round10(a) + round10(b) + round10(c))

def round10(num):
  if num%10>=5:

    print(num/10*10 + 10)
  else:
    print( num/10*10)