import requests
from bs4 import BeautifulSoup
import smtplib
import time

URL = "https://www.amazon.de/Samsung-LC27F396FHUXEN-Monitor-Reaktionszeit-schwarz/dp/B01BCF09R0/ref=sr_1_3?__mk_de_DE=%C3%85M%C3%85%C5%BD%C3%95%C3%91&keywords=computer+monitor&qid=1579185889&sr=8-3"

headers = {"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36"}
def check_price():
    page = requests.get(URL, headers=headers)

    soup = BeautifulSoup(page.content, "html.parser")

    title = soup.find(id="productTitle").get_text()
    #print(title)

    price = soup.find(id="priceblock_ourprice").get_text()
    converted_price = float(price[0:3])

    if(converted_price < 100):
        send_email()
    print(converted_price)
    print(title.strip())


def send_email():
    server = smtplib.SMTP("smtp.gmail.com", 587)
    server.ehlo()
    server.starttls()
    server.ehlo()

    server.login("jakwai01@gmail.com", "mzcbtuoffveylyey")
    subject = "Price fell down!"
    body = "Check this link: URL = https://www.amazon.de/Samsung-LC27F396FHUXEN-Monitor-Reaktionszeit-schwarz/dp/B01BCF09R0/ref=sr_1_3?__mk_de_DE=%C3%85M%C3%85%C5%BD%C3%95%C3%91&keywords=computer+monitor&qid=1579185889&sr=8-3"

    msg =f"Subject: {subject}\n\n{body}"

    server.sendmail(
        "jakwai01@gmail.com",
        "jakob.waibel@web.de",
        msg
    )

    print("EMAIL SENT")

    server.quit()

while(True):
    check_price()
    time.sleep(10)