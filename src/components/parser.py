import re

class ParseTree():
    def __init__(self, left=None, root=None, right=None):
        self.left = left 
        self.root = root
        self.right = right
    
class StackNode():
    def __init__(self, value):
        self.value = value 
        self.next = None

class Stack():
    def __init__(self):
        self.head = StackNode("head")
        self.size = 0

    def __str__(self):
        cur = self.head.next
        out = ""
        while cur:
            out += str(cur.value) + "->"
            cur = cur.next
        return out[:-3]

    def getSize(self):
        return self.size

    def isEmpty(self):
        return self.size == 0

    def peek(self):
        if self.isEmpty():
            raise Exception("Peeking from an empty stack")
        return self.head.next.value

    def push(self, value):
        node = StackNode(value)
        node.next = self.head.next 
        self.head.next = node
        self.size += 1

    def pop(self):
        if self.isEmpty():
            raise Exceptions("Popping from an empty stack")
        remove = self.head.next
        self.head.next = self.head.next.next
        self.size -= 1
        return remove.value

class Parser():
    def __init__(self: object):
        pass
    

    def build_parse_tree(self: object, tokens: list[str]):
        stack = Stack()
        root = ParseTree() 
        stack.push(root)
        current_node = root 

        for token in tokens: 
            # print( 4 + 4 - 2 );
            if token == "(":
                current_node.left = ParseTree()
                stack.push(current_node)
                current_node = current_node.left
            if token in ["+", "-"]:
                current_node.root = token
                current_node.right = ParseTree()
                stack.push(current_node)
                current_node = current_node.right
            if token.isdigit():
                current_node.root = int(token)
                parent = stack.pop()
                current_node = parent
            if token == ")":
                stack.pop()

        return current_node
