import tkinter as tk

wincounter = 0
loosecounter = 0

def wincounter_label(label):
    wincounter = 0

def wincount():
    global wincounter
    wincounter += 1
    label.config(text=str(wincounter))

def loosecounter_label(label):
    loosecounter = 0

def loosecount():
    global loosecounter
    loosecounter += 1
    label2.config(text=str(loosecounter))

root = tk.Tk()
root.title("Games")
label = tk.Label(root, fg="dark green")
label.pack(side=tk.TOP)
label2 = tk.Label(root, fg="dark red")
label2.pack(side=tk.BOTTOM)
wincounter_label(label)
loosecounter_label(label2)
button = tk.Button(root, text='Win', width=25, command=wincount)
button.pack()
button2 = tk.Button(root, text='Loose', width=25, command=loosecount)
button2.pack()
root.mainloop()