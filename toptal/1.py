#env python3
# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")
import re
def solution(N, S):
    # write your code in Python 3.6
    seats = []
    for i in range(1, N+1):
        seats.append([0]*10)
        for seat in re.findall(r"{}[A-K]".format(i), S):
            seats[i-1][ord(seat[-1])-65] = 1

    print(seats)
    # count continuos four seats
    nfour_seats = 0
    eseats = [0]*4
    for row in seats:
        count = 0
        if row[1:5] == eseats:
            count += 1
        if row[5:9] == eseats:
            count += 1
        if 0 == count:
            if row[3:7] == eseats:
                count +=1
        nfour_seats += count

    return nfour_seats

if "__main__" == __name__:
    print(solution(2, '1A 2F 1C'))
    print(solution(2, '1D 2F 1C'))
