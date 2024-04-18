class A:
    pass

class B(A):
    pass

class C(A):
    pass

class D(B, C):
    pass

class E(C, B):
    pass

class F(D, E):
    pass

f = F()

"""
The C3 linearization algorithm is a method used in Python to determine the method resolution order (MRO) in the presence of multiple inheritance. It ensures that classes are searched for methods and attributes in a consistent and predictable order.

Here's a simple explanation of the C3 linearization algorithm with an example:

Suppose we have the following class hierarchy:

```python
class A:
    pass

class B(A):
    pass

class C(A):
    pass

class D(B, C):
    pass
```

To determine the method resolution order (MRO) for class `D`, the C3 linearization algorithm works as follows:

1. Start with the class itself: `D`
2. Consider the inheritance list from left to right (`B` before `C`):
    - `D(B, C)`: The MRO of `B` should be before `C`
3. For each class in the inheritance list:
    - If the class doesn't have any unresolved base classes, add it to the MRO list.
    - If the class has unresolved base classes, recursively resolve their MRO first.
4. Merge the MRO lists of the base classes while preserving the order and ensuring that each class appears only once.
5. Add the current class to the beginning of the merged MRO list.

Using the C3 linearization algorithm, the MRO for class `D` in our example would be:

```
MRO(D) = [D] + merge(MRO(B), MRO(C), [B, C])
       = [D] + merge([B, A], [C, A], [B, C])
       = [D, B] + merge([A], [C, A], [C])
       = [D, B, C] + merge([A], [A])
       = [D, B, C, A]
```

So, the method resolution order (MRO) for class `D` is `[D, B, C, A]`. This means that when you access an attribute or method on an instance of class `D`, Python will first look in `D`, then `B`, then `C`, and finally `A`, ensuring a consistent and predictable order for method and attribute resolution.

"""
