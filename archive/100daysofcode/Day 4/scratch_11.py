#Datetimes

import datetime

today = datetime.date.today()

print(today)

birthday = datetime.date(2001, 5, 24)
print(birthday)

days_since_birth = (today - birthday).days
print(days_since_birth)

tdelta = datetime.timedelta(days=10)
print(today + tdelta)

print(today.day)
print(today.weekday())

print(datetime.time(7, 2, 20, 15))

hour_delta = datetime.timedelta(hours=10)
print(datetime.datetime.now() + hour_delta)

datetime.datetime.now()