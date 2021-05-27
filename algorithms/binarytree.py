class BTree:
    def __init__(self, value, left_child=None, left_right=None):
        self.value = value
        self.left_child = left_child
        self.right_child = left_right

tree = BTree(1, BTree(2, BTree(4), BTree(5)), BTree(3, BTree(6), BTree(7)))

def invertTree(tree):
    if tree:
        left_child = tree.left_child
        right_child = tree.right_child
        tree.left_child = right_child
        tree.right_child = left_child
        invertTree(tree.left_child)
        invertTree(tree.right_child)

print(tree.value)
print(tree.left_child.value)
print(tree.right_child.value)
print(tree.left_child.left_child.value)
print(tree.left_child.right_child.value)
print(tree.right_child.left_child.value)
print(tree.right_child.right_child.value)

invertTree(tree)

print(tree.value)
print(tree.left_child.value)
print(tree.right_child.value)
print(tree.left_child.left_child.value)
print(tree.left_child.right_child.value)
print(tree.right_child.left_child.value)
print(tree.right_child.right_child.value)

