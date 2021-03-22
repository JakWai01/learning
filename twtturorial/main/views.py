from django.shortcuts import render
from django.http import HttpResponse, HttpResponseRedirect
from .models import TodoList, Item
from .forms import CreateNewList

# Create your views here.
def index(request, id):
    ls = TodoList.objects.get(id=id)

    if ls in request.user.todolist.all():

        if request.method == "POST":
            print(request.POST)
            if request.POST.get("save"):
                for item in ls.item_set.all():
                    if request.POST.get("c" + str(item.id)) == "clicked":
                        item.complete = True
                    else: 
                        item.complete = False
                    item.save()

            elif request.POST.get("newItem"):
                txt = request.POST.get("new")

                if len(txt) > 2:
                    ls.item_set.create(text=txt, complete=False)
                else: 
                    print("invalid")

        return render(request, "main/list.html", {'ls': ls})
    return render(request, "main/view.html", {})

def home(request):
    return render(request, "main/home.html", {})

def create(request):

    if request.method == "POST":
        form = CreateNewList(request.POST)

        if form.is_valid():
            n = form.cleaned_data["name"]
            t = TodoList(name=n)
            t.save()
            request.user.todoList_set.create(name=t)
            

        return HttpResponseRedirect("/%i" %t.id)
    else:
        form = CreateNewList()
    return render(request, "main/create.html", {"form":form})

def view(request):
    return render(request, "main/view.html", {})