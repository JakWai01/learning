from django.db import models
from django.urls import reverse

# Create your models here.
class Album(models.Model):
    artist = models.CharField(max_length=255)
    album_title = models.CharField(max_length=511)
    genre = models.CharField(max_length=127)
    album_logo = models.FileField()

    def get_absolute_url(self):
        return reverse("music:detail", kwargs={"pk": self.pk})
    
    def __str__(self):
        return self.album_title + " - " + self.artist

class Song(models.Model):
    album = models.ForeignKey(Album, on_delete=models.CASCADE)
    file_type = models.CharField(max_length=15)
    song_title = models.CharField(max_length=255)
    is_favorite = models.BooleanField(default=False)

    def __str__(self):
        return self.song_title