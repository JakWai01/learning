from file import read_file
from lexer import Lexer

text_input = read_file("../../samples/code.sucks")

lexer = Lexer()
tokens = lexer.lex(text_input)
print(tokens)

