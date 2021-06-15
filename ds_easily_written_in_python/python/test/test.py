import os
import sys


print(os.path.dirname(__file__))

print(os.path.abspath(os.path.dirname(__file__)))

print(sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(__file__)))))))

print((os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(__file__)))))))