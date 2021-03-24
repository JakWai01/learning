from django import forms
from django.forms import TextInput

class CreateNewList(forms.Form):
    text = forms.CharField(label="Text", max_length=200, widget=forms.TextInput(attrs={'class': 'focus:outline-none focus:ring focus:border-blue-300 rounded border-2'}))
    check = forms.BooleanField(required=False)