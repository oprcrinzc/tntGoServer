i = input(": ")
s = ""
for j in i:
  f = ""
  if ord(j) >= 97:
    f = chr(ord(j)-32)
  else:
    f = chr(ord(j)+32)
  s+=f
print(s)
