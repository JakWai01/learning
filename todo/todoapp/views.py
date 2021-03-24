from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect
from .models import Todo
from .forms import CreateNewList

# Create your views here.


def index(request):
    todos = Todo.objects.all()
    return render(request, "todoapp/todos.html", {"todos": todos})


def create(request):

    if request.method == "POST":
        form = CreateNewList(request.POST)

        if form.is_valid():
            n = form.cleaned_data["text"]
            t = Todo(text=n)
            t.save()
            #request.user.todo_set.create(text=t)

        return HttpResponseRedirect("/")
    else:
        form = CreateNewList()
        return render(request, "todoapp/create.html", {"form": form})
