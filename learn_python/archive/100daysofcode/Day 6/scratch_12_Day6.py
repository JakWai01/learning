import requests
from bs4 import BeautifulSoup
page = requests.get("https://www.wetter.de/deutschland/wetter-stuttgart-neugereut-23118224193.html")
soup = BeautifulSoup(page.content, "html.parser")
print(soup.find_all("a"))
4:11:54