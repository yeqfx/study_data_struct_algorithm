import os
import sys

sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))))

from dataStruct import ArrayLinearQueue

q = ArrayLinearQueue.ArrayLinearQueue()
for i in range(8):
    q.enqueue(i)
print(q)
for i in range(5):
    q.dequeue()
print(q)
for i in range(8, 14):
    q.enqueue(i)
print(q)