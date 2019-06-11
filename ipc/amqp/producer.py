from __future__ import absolute_import, unicode_literals

import datetime

from kombu import Connection, Exchange, Queue
from kombu.simple import SimpleQueue
from amqp_buffer import AmqpBuffer

with Connection('amqp://guest:guest@192.168.157.128:5672/') as conn:
    #simple_queue = AmqpBuffer(conn, 'simple_queue')
    for i in range(2):
        queue_name = "que"+str(i)
        exchange = Exchange(name="event1", type="direct")
        queue = Queue(queue_name, exchange=exchange, routing_key=queue_name)
        simple_queue = SimpleQueue(conn, queue)
        message = 'helloworld, sent at {0}'.format(datetime.datetime.today())
        simple_queue.put(message)
        print(queue_name,'Sent: {0}'.format(message))
        simple_queue.close()
