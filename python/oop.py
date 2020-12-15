class Computer:

    def __init__(self):
        self.__maxprice__ = 900
        self._minprice = 100

    def sell(self):
        print("Selling Price: {}".format(self.__maxprice__))
        #print("Selling  min Price: {}".format(self._minprice))

    def setMaxPrice(self, price):
        self.__maxprice__ = price

c = Computer()
c.sell()

# change the price
c.__maxprice__ = 1000
c._minprice = 80
c.sell()

# using setter function
c.setMaxPrice(2000)
c.sell()

print(dir(c))