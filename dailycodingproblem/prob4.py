"""
This problem was asked by Facebook.

Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.

For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.

You can assume that the messages are decodable. For example, '001' is not allowed.
"""
def num_way(data):

    def helper(encoded, k):
        print(k)
        if k == 0:
            return 1
        if encoded[len(encoded)-k] == '0':
            return 0

        result = helper(encoded, k-1)

        if k >=2 and encoded[k:k+2] != ""  and int(encoded[k:k+2]) <=26:
            result += helper(encoded, k-2)

        return result

    return helper(data, len(data))

if __name__ == "__main__":
    print(num_way("1231"))
