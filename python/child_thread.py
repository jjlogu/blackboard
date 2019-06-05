import threading
import time

class parentThread(threading.Thread):
    #def __init__(self, val):
    i=2
    def run(self):
        while self.i >0:
            if self.i ==2:
                t=childThread()
                t.start()
                #t.join()
            print("Parent", self.i)
            time.sleep(1)
            self.i -= 1
        print("Parent exit")
    def kill(self):
        print("Parent kill")
        self.i=3

class childThread(threading.Thread):
    #def __init__(self, val):
    j=3
    def run(self):
        while self.j > 0:
            print("Child", self.j)
            time.sleep(2)
            self.j -= 1
        print("Child exit")

if __name__ == "__main__":
    p = parentThread()
    p.start()
    while True:
        print(p.is_alive())
        if not p.is_alive():
            #p.kill()
            p.start()
            break
        time.sleep(0.5)
    #p.join()
    print("Program exit")
