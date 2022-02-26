from file import read_file
from lexer import Lexer
from parser import Parser

text_input = read_file("../../samples/code.sucks")

lexer = Lexer()
tokens = lexer.lex(text_input)

parser = Parser()
parse_tree = parser.build_parse_tree(tokens)

print(parser.eval(parse_tree))