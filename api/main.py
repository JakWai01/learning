import requests 

# Getting random Taylor Swift lines 
response = requests.get("https://api.taylor.rest/")

data = response.json()

print(data["quote"])
