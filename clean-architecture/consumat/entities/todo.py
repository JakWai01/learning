class Todo:
    def __init__(self, task, done):
        self.task = task
        self.done = done

    def __str__(self):
        return self.task
