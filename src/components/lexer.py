class Lexer():
    _symbols = ['(', ')', '+' '-', ';']
    _white_space = ' '
    _lexeme = ''

    def lex(self: object, text_input: str) -> list[str]:
        tokens = []

        for i, char in enumerate(text_input):
            if char != self._white_space:
                self._lexeme += char
            if (i+1 < len(text_input)):
                if text_input[i+1] == self._white_space or text_input[i+1] in self._symbols:
                    if self._lexeme != '':
                        tokens.append(self._lexeme.replace('\n', '<newline>'))
                        self._lexeme = ''
        return tokens