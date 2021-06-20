class ArrayLinearQueue:
    
    def __init__(self):
        self.items = []

    def __str__(self):
        return str(self.items)

    def isEmpty(self):
        return len(self.items) == 0

    def enqueue(self, item):
        self.items.append(item)

    def dequeue(self):
        if not self.isEmpty():
            return self.items.pop(0)

    def peek(self):
        if not self.isEmpty():
            return self.items[0]

    def size(self):
        return len(self.items)

    def clear(self):
        self.items = []
