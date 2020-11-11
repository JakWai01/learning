import calendar

#Gibt die Wochentage mit x Buchstaben an
print(calendar.weekheader(10))
print()
#Gibt uns den ersten Wochentag aus
print(calendar.firstweekday())
print()
#Gibt den gewünschten Monat aus
#Buchstabenanzahl vor Monat ändern?? In Dokumentation nachschauen
print(calendar.month(2020, 1))
print(calendar.monthcalendar(2020, 1))
print(calendar.calendar(2020))
day_of_the_week= calendar.weekday(2020, 1, 4)
print(day_of_the_week)
#Prüfen ob ein Jahr ein Schaltjahr ist
is_leap = calendar.isleap(2020)
print(is_leap)
#Wie viele SchaltTAGE gab es von 2000 bis 2003, da exklusive Klammer
how_many_leap_days = calendar.leapdays(2000,2004)
print(how_many_leap_days)