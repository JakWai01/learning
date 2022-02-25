import re
from ast import Sub, Sum, BinaryOp, Number

class ParseTree():
    def __init__(self, left=None, root=None, right=None):
        self.left = left 
        self.root = root
        self.right = right

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
            if token == "(":
                # 1 create tree on left side and go down
                current_node.left = ParseTree()
                stack.push(current_node)
                current_node = current_node.left
            if token in ["+", "-"]:
                # 3 set current root to "+", create tree on left and go there
                # 5 set current root to "-", create tree on left and go there
                current_node.root = token
                current_node.right = ParseTree()
                stack.push(current_node)
                current_node = current_node.right
            if token.isdigit():
                # 2 set value of tree to "4" and go to parent
                # 4 set value of tree to "4" and go to parent
                # 6 set value of tree to "2" and go to parent
                current_node.root = int(token)
                parent = stack.pop()
                current_node = parent
            if token == ")":
                # 7 move to parent
                current_node = stack.pop()

        return current_node

    def eval(self, tree):
        if tree.root == "+":
            return Sum(BinaryOp(self.eval(tree.left), self.eval(tree.right))).eval()
        if tree.root == "-":
            return Sub(BinaryOp(self.eval(tree.left), self.eval(tree.right))).eval()
        if type(tree.root) == type(4):
            return int(tree.root)



