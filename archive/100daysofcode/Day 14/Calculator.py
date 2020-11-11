print("Enter number:")
first_digit = int(input())
print("Enter second number:")
second_digit = int(input())
print("Choose + | - | * | /:")
operator = input()

if operator == "+":
    result = first_digit + second_digit
    print(result)
if operator == "-":
    result = first_digit - second_digit
    print(result)
if operator == "/":
    result = first_digit / second_digit
    print(result)
if operator == "*":
    result = first_digit * second_digit
    print(result)

