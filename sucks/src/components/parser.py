import re
import ast

class ParseTree():
    def __init__(self, left=None, root=None, right=None, direct_child=None):
        self.left = left 
        self.root = root
        self.right = right
        self.direct_child = direct_child

# It's also possible using a linked list instead of implementing it with a list.
# For the sake of brevity, we use a list. 
class Stack:
    def __init__(self):
        self.items = []

    def isEmpty(self):
        return self.items == []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        return self.items.pop()
    
    def peek(self):
        return self.items[len(self.items)-1]
    
    def size(self):
        return len(self.items)

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
            if token == "print":
                current_node.root = token
                current_node.direct_child = ParseTree() 
                current_node = current_node.direct_child

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
                current_node.root = ast.Number(token)
                parent = stack.pop()
                current_node = parent
            if token == ")":
                current_node = stack.pop()

        return current_node

    def eval(self, tree):
        if tree.root == "print":
            print(self.eval(tree.direct_child)) 
        if tree.root == "+":
            return self.eval(tree.left) + self.eval(tree.right)
        if tree.root == "-":
            return self.eval(tree.left) - self.eval(tree.right)
        if type(tree.root) == ast.Number:
            return tree.root.eval()



