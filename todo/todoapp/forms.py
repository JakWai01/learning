from django import forms
from django.forms import TextInput

class CreateNewList(forms.Form):
    text = forms.CharField(label="", max_length=200, widget=forms.TextInput(attrs={'class': 'focus:outline-none focus:ring focus:border-blue-300 rounded border-2 mb-2 w-max'}))