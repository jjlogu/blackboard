

# 5*4*3*2*1
# 0! = 1

def factorial(n):


    if n == 1 or n == 0:
        return 1
    if n< 0:
        return -1
    return n * factorial(n-1) 


print(factorial(6))
