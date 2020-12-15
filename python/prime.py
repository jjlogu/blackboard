def is_prime(n):

    for i in range(2,n/2):
        if n%i == 0:
            return False
    return True
def _prime_number(n):
    # prime number for a given range`
    # 1, 2, 3, 4, 5 , 6 ,7
    p_numbers = []
    for i in range(1, n+1):
        if i in [1 ,2 ,3]:
            p_numbers.append(i)
        elif is_prime(i):
            p_numbers.append(i)
    return p_numbers

def prime_number(rang):
    p = []
    count = 0 # num_of_p_numbers
    n = 1
    while True:
        if is_prime(n):
            p.append(n)
            count += 1
        if count == rang:
            break
        n += 1

    return p

if __name__ == "__main__":
    print(prime_number(10))
