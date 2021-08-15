class ListTodos():
    def __init__(self, repo):
        self.repo = repo
    def __repr__(self):
        return str(self.repo.list())

