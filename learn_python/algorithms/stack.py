class StackOverflowError(BaseException):
    pass

class Stack:
    
    def __init__(self, limit: int = 10):
        self.stack = []
        self.limit = limit

    def __bool__(self) -> bool:
        return bool(self.stack)

    def __str__(self) -> str:
        return str(self.stack)

    def push(self, data):
        """Push an element to the top of the stack."""
        if len(self.stack) >= self.limit:
            raise StackOverflowError
        self.stack.append(data)

    def pop(self):
        """Pop an element to the top of the stack."""
        return self.stack.pop()

    def peek(self):
        """Peek at the top-most element of the stack."""
        return self.stack[-1]

    def is_empty(self) -> bool:
        """Check if a stack is empty."""
        return not bool(self.stack)

    def is_full(self) -> bool:
        return self.size() == self.limit

    def size(self) -> int:
        """Return the size of the stack."""
        return len(self.stack)

    def __contains__(self, item) -> bool:
        """Check if item is in stack"""
        return item in self.stack

stack = Stack()

stack.push(1)
stack.push(2)
stack.push(4)
stack.push(5)
stack.push(6)

stack.pop()

print(stack)
