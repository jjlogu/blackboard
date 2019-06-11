from __future__ import absolute_import, unicode_literals

import datetime

from kombu import Connection
from amqp_buffer import AmqpBuffer

with Connection('amqp://guest:guest@192.168.157.128:5672/') as conn:
    simple_queue = AmqpBuffer(conn, 'simple_queue')
    message = 'helloworld, sent at {0}'.format(datetime.datetime.today())
    simple_queue.put(message)
    print('Sent: {0}'.format(message))
    simple_queue.close()
