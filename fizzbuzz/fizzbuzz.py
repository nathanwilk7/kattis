from sys import stdin

x,y,n = stdin.readline().split(' ')
x = int(x)
y = int(y)
n = int(n)
is_div_by_x = set()
x_test = x
while x_test <= n:
    is_div_by_x.add(x_test)
    x_test += x
is_div_by_y = set()
y_test = y
while y_test <= n:
    is_div_by_y.add(y_test)
    y_test += y
for i in range(1, n+1):
    if i in is_div_by_x and i in is_div_by_y:
        print('FizzBuzz')
    elif i in is_div_by_x:
        print('Fizz')
    elif i in is_div_by_y:
        print('Buzz')
    else:
        print(i)