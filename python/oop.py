class Computer:
    class_var = 10

    def __init__(self):
        self.__maxprice__ = 900
        self._minprice = 100

    def sell(self):
        print("Selling Price: {}".format(self.__maxprice__))
        #print("Selling  min Price: {}".format(self._minprice))

    def setMaxPrice(self, price):
        self.__maxprice__ = price

print("Computer.class_var",Computer.class_var)
c = Computer()
c.sell()
c.class_var = 100
print("Computer.class_var",Computer.class_var)
print("c.class_var",c.class_var)

# change the price
c.__maxprice__ = 1000
c._minprice = 80
c.sell()

# using setter function
c.setMaxPrice(2000)
c.sell()

print(dir(c))
