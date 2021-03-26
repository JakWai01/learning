class Todo:
    def __init__(self, task, done):
        self.task = task
        self.done = done

    def __repr__(self):
        return self.task
