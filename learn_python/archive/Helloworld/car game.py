interface = ""
stop = "stops car"
start = True

while interface != "exit":

    interface = input(">")
    if interface == "help":
        print("start = starts the car")
        print("stop = stops the car")
        print("exit = exit app")

    if interface == "start":

        print("The car has been started")
        stop = input(">")
        if stop == "stop" and start == bool(True):
             print("The car has been stopped")
        if stop == "start" and start == bool(True):
             print("You can't start your car twice")

    if interface == "stop":
        print(" You cant stop your car if it hasn't been started yet")

    if interface == exit:
        print("bye bye")
        break