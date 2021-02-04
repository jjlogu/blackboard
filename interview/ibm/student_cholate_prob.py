# find repeated letters in a given strings

# str= input()

# res = {
#     a:1,

# }


# 5 students sitting in a circular fashion

# -> s1 -> s2 -> s3 -> s4 -> s5

# input: number students
# student to begin with: int
# number of choclate

# return last person to hold choclate

num_students = int(input())
begin = int(input())
num_choc = int(input())

# while True:
#     current_person = begin
#     num_choc -=1
#     if num_choc == 0:
#         print(current_person)
#         break
#     if current_person == num_students:
#         current_person = 1
#     else:
#         current_person += 1

last_person = ((num_choc%num_students)+begin-1)%num_students

print(last_person)
