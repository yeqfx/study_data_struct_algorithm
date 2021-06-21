import os
import sys

sys.path.append(os.path.dirname(os.path.abspath(os.path.dirname(os.path.abspath(os.path.dirname(__file__))))))

from dataStruct import ArrayCircularQueue

def isValidPos(x, y):
    if x < 0 or y < 0 or x >=MAZE_SIZE or y >= MAZE_SIZE:
        return False
    else:
        return map[y][x] == '0' or map[y][x] == 'x'

def BFS():
    que = ArrayCircularQueue.ArrayCircularQueue()
    que.enqueue((0, 1))
    print('BFS :')

    while not que.isEmpty():
        here = que.dequeue()
        print(here, end='->')
        x, y = here
        if (map[y][x] == 'x'):
            return True
        else:
            map[y][x] = '.'
            if isValidPos(x, y-1): que.enqueue((x, y-1))    # 상
            if isValidPos(x, y+1): que.enqueue((x, y+1))    # 하            
            if isValidPos(x-1, y): que.enqueue((x-1, y))    # 좌
            if isValidPos(x+1, y): que.enqueue((x+1, y))    # 우

    return False


map = [['1', '1', '1', '1', '1', '1'],
       ['e', '0', '1', '0', '0', '1'],
       ['1', '0', '0', '0', '0', '1'],
       ['1', '0', '1', '0', '1', '1'],
       ['1', '0', '1', '0', '0', 'x'],
       ['1', '1', '1', '1', '1', '1']]

MAZE_SIZE = 6

result = BFS()
if result:
    print(' --> 미로탐색 성공')
else:
    print(' --> 미로탐색 실패')