from sys import stdin

count = stdin.readline()
count = int(count.strip())
for i in range(1,count+1):
    print('{} Abracadabra'.format(i))
