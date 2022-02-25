from file import read_file
from lexer import Lexer
from parser import Parser

text_input = read_file("../../samples/code.sucks")

lexer = Lexer()
tokens = lexer.lex(text_input)
print(tokens)

parser = Parser()
parse_tree = parser.build_parse_tree(tokens)

print("Root", parse_tree.root)
print("Left root", parse_tree.left.root)
print("Right root", parse_tree.right.root)
print("Left root left child", parse_tree.left.left.root)
print("Left root right child", parse_tree.left.right.root)
print(parser.eval(parse_tree))