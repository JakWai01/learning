#Email address text scraper

import re
#Der Punkt in einer Email macht was aus
text = "A random string yourname@web.de, my.name@web.de, jakob.waibel@web.de"

pattern = re.compile("A random string")
#sucht nur den ersten Match danach hört es auf. Ein Beweis reicht damit es als gefunden gilt
pattern = re.compile("[a-zA-Z0-9]+\.[a-zA-Z]+\@[a-zA-Z0-9]+\.[a-zA-Z]+")

result = pattern.findall(text)

print(result)
