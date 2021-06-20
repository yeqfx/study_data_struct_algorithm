import os
import sys

sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))))

from dataStruct import ArrayCircularQueue

q = ArrayCircularQueue.ArrayCircularQueue()
for i in range(8):
    q.enqueue(i)
q.display()
for i in range(5):
    q.dequeue()
q.display()
for i in range(8, 14):
    q.enqueue(i)
q.display()