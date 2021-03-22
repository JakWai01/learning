from django.db import models

# Create your models here.
class TodoList(models.Model):
    name = models.CharField(max_length=255)

    def __str__(self):
        return self.name

class Item(models.Model):
    todoList = models.ForeignKey(TodoList, on_delete=models.CASCADE)
    text = models.CharField(max_length=255)
    complete = models.BooleanField(default=False)

    def __str__(self):
        return self.text