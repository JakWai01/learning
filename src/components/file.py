def read_file(path: str) -> str:
    with open(path) as f: 
        text_input = f.read()
    return text_input