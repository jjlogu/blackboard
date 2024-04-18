# Generator function to generate squares of numbers
def squares(n):
    for i in range(n):
        yield i ** 2

# Generator function to filter even numbers
def filter_even(nums):
    print(f"nums: {nums}")
    for num in nums:
        print(f"num: {num}")
        if num % 2 == 0:
            yield num

# Generator function to take the first n items
def take(n, iterable):
    for i, item in enumerate(iterable):
        if i >= n:
            break
        yield item

# Define the pipeline
data = range(10)
pipeline = take(5, filter_even(squares(10)))

# Consume the pipeline
for item in pipeline:
    print(item)

