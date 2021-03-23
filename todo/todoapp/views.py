from django.shortcuts import render
from django.http import HttpResponse
from .models import Todo
# Create your views here.
def index(request):
    return HttpResponse(str(Todo.objects.get(id=1)))