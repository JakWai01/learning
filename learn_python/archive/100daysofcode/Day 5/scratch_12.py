import pandas as pd
import requests
from bs4 import BeautifulSoup
response = requests.get("https://www.wetter.de/deutschland/wetter-stuttgart-neugereut-23118224193.html")
soup = BeautifulSoup(response.content, "html.parser")
#print(soup.find_all("a"))
#Der untenstehende Befehl prüft, ob alles geklappt hat und die Website erfolgreich überprüft wurden
#(200 = OK, 404 = ERROR), könnte man auch direkt ausgeben lassen
response.status_code
#Nun den Code in verständliche Worte konvertieren
#if response.status_code == 200:
    #print("Success!")
#if response.status_code == 404:
    #print("Not Found!")
#print(soup)
week = soup.find(id="item-20200114")


items = week.find_all(class_="forecast-column column-1 wt-border-radius-6")

#print(items[0])
print(items[0].find(class_="forecast-hour-statement wt-font-semibold").get_text())
print(items[0].find(class_="forecast-column-condition").get_text())
print(items[0].find(class_="forecast-text-temperature wt-font-light").get_text())

period_names = [item.find(class_="forecast-hour-statement wt-font-semibold").get_text() for item in items]
short_descriptions = [item.find(class_="forecast-column-condition").get_text() for item in items]
temperatures = [item.find(class_="forecast-text-temperature wt-font-light").get_text() for item in items]

print(period_names)
print(short_descriptions)
print(temperatures)

weather_stuff = pd.DataFrame(
    {"period": period_names, "short_desctiptions" : short_descriptions, "temperatures" : temperatures})
print(weather_stuff)

#weather_stuff.to_csv("weather.csv")
#Man könnte noch zusätzlich die jeweilige Uhrzeit und das Datum anzeigen lassen, da die Dokumente sonst etwas unorganisiert sind