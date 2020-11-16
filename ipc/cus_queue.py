import threading
import queue
import logging
import time
logging.basicConfig(
    level=logging.DEBUG,
    format='(%(threadName)-10s) %(message)s',
)

num_worker_threads=3 
def worker():
    while True:
        item = q.get()
        if item is None:
            break
        time.sleep(0.01)
        do_work(item)
        q.task_done()

def do_work(item):
    logging.info(item)

q = queue.Queue()
threads = []
for i in range(num_worker_threads):
    t = threading.Thread(name="thread"+str(i),target=worker)
    t.start()
    threads.append(t)


for item in range(num_worker_threads):
    q.put([123,i])

# block until all tasks are done
q.join()

# stop workers
for i in range(num_worker_threads):
    q.put(None)
for t in threads:
    t.join()
