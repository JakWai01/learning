print("Bitte gib hier deinen Namen ein: ")
name = input()
print("Print start to begin your journey: \n[START]")
start = input()
if start == "start".lower():
    print("Wenn die Menschheit einmal so weit ist Simulationen einer Menschheit beziehungsweise eines Universums mit Hilfe des Computers zu erschaffen,\nso ist die wahrscheinlichkeit, dass wir ebenfalls nur eine Simulation sind nahezu 100%. Diese Worte hörts du einen genia"\n[INSPECT] [REJECT]")
else: print("Please enter start to begin.")

inspect_browser = input()
if inspect_browser == "INSPECT".lower():
    exit()
elif inspect_browser == "REJECT".lower():
    print("Du schließt deinen Browser mit leichter verwirrung und beschließt heute einmal früher ins Bett zu gehen. Gute Nacht!")
    exit("Du scheinst nicht bereit für dieses Abenteuer zu sein!")
else: print("Please enter a valid input " + name)

