from __future__ import absolute_import, unicode_literals

from kombu import Connection, entity
from kombu.five import Empty
from amqp_buffer import AmqpBuffer
# import threading

exchange = entity.Exchange(name="target_event", type="direct")

#: Create connection
#: If hostname, userid, password and virtual_host is not specified
#: the values below are the default, but listed here so it can
#: be easily changed.
with Connection('amqp://guest:guest@192.168.157.128:5672//') as conn:
    #: SimpleQueue mimics the interface of the Python Queue module.
    #: First argument can either be a queue name or a kombu.Queue object.
    #: If a name, then the queue will be declared with the name as the queue
    #: name, exchange name and routing key.
    with AmqpBuffer(conn, 'simple_queue') as queue:
        messages=[]
        messages.append(queue.get(block=True))
        messages[-1].ack()

        # Check if there is anymore messages
        try:
            while True:
                messages.append(queue.get(block=True, timeout=1))
                messages[-1].ack()
        except Empty:
            pass

        for message in messages:
            print(message.payload)

####
#: If you don't use the with statement then you must aways
# remember to close objects after use:
#   queue.close()
#   connection.close()

"""
class DT(threading.Thread):
    def __init__(self, connection, exchange, queue_name):
        self.connection = connection
        self.buffer = AmqpBuffer(self.connection, entity.Queue(queue_name, exchange=exchange))
"""