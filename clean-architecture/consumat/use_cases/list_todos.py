class ListTodos():
    def __init__(self, repo):
        self.repo = repo
    
    def __str__(self):
        return repo.list()
