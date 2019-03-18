#!/bin/python3
'''
a ->1
b ->2
.
.
.
z->26

input:
    data: "12" 
output:
    "ab","l"

input:
    data: "01"
output:
    ""

'''

import string

s2i = {}
for letter in string.ascii_lowercase:
    s2i["{}".format(ord(letter)-96)]=letter

def nways_num(data):
    result = [""]
    for i in data:
        if not s2i.get(i):
            return []
        result[0] += s2i[i]
    if s2i.get(data):
        result.append(s2i[data])
    return result

data = input()
print(nways_num(data))

