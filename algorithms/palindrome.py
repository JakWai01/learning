palindrome = "lagerregal"

def is_palindrome(s):
    print(s)
    if len(s) < 2:
        return True
    else:
        if not s[0] == s[len(s)-1]:
            return False
        else:
            return is_palindrome(s[1:len(s)-1])

print(is_palindrome(palindrome))
