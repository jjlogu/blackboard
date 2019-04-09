"""
This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'
"""
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
    def __repr__(self):
        return '{}(value={!r}, left={!r}, right={!r})'.format(self.__class__.__name__, self.val, self.left, self.right)

def inorder(root):
    if not root:
        return None
    inorder(root.left)
    print(root.val)
    inorder(root.right)

def preorder(root):
    if not root:
        return None
    print(root.val)
    preorder(root.left)
    preorder(root.right)


def serialize1(root, prefix=""):
    if not root:
        return ""
    return "{} {}{}".format(root.val, serialize(root.left), serialize(root.right))

def serialize(root, output):
    if not root:
        output.append('#')
        return
    output.append(root.val)
    serialize(root.left, output)
    serialize(root.right, output)
    return output

def deserialize(string, root=None):
    if not string:
        return ""
    
    val , string = string.split(" ", 1)
    if not root:
        pass
        # root = Node(valkk

if "__main__" == __name__:
    node = Node('root', Node('left', Node('left.left')), Node('right',right=Node('right.right')))
    # preorder(node)
    print(serialize(node, []))
    # print(node)

