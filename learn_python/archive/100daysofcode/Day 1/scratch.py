import turtle

a=5
print(a)

b="hello"
print(b)
qazi_turtle = turtle.Turtle()
def square():

    # This is a square
    qazi_turtle.forward(100)
    qazi_turtle.right(90)
    qazi_turtle.forward(100)
    qazi_turtle.right(90)
    qazi_turtle.forward(100)
    qazi_turtle.right(90)
    qazi_turtle.forward(100)

square()
qazi_turtle.forward(200)
square()
#Unnötig, aber dann ist der Kreis geschlossen
qazi_turtle.right(90)
qazi_turtle.forward(100)