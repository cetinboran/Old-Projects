password = input("Enter your password: ")

RATE_LEN = 0.4
RATE_NUMBER = 0.1
RATE_SPECIAL = 0.4
RATE_UPPER = 0.1

upperChCount = 0
specialChCount = 0
numberChCount = 0

for ch in password:
    if ch != ch.lower():
        upperChCount += 1
    elif ord(ch) < 48 or ord(ch) > 57 and ord(ch) < 65 or ord(ch) > 90 and ord(ch) < 97 or ord(ch) > 122:
        specialChCount +=1
    elif ord(ch) >= 48 and ord(ch) <= 57:
        numberChCount +=1


def LEN(password):
    strength_LEN = 0
    
    if len(password) < 4:
        strength_LEN += 0
    elif len(password) < 6:
        strength_LEN += 1
    elif len(password) < 8:
        strength_LEN += 2
    else:
        strength_LEN += 3
    
    return strength_LEN * RATE_LEN

def NUMBER(password):
    strength_NUMBER = 0

    if numberChCount == 0: return 0

    if numberChCount < 2:
        strength_NUMBER += 1
    elif numberChCount < 4:
        strength_NUMBER += 2
    else:
        strength_NUMBER += 3

    return strength_NUMBER * RATE_NUMBER

def SPECIAL(password):
    strength_SPECIAL = 0

    if specialChCount == 0: return 0

    if specialChCount < 2:
        strength_SPECIAL += 1
    elif specialChCount < 4:
        strength_SPECIAL += 2
    else:
        strength_SPECIAL += 3

    return strength_SPECIAL * RATE_SPECIAL

def UPPER(password):
    strength_UPPER = 0

    if upperChCount == 0: return 0

    if upperChCount < 2:
        strength_UPPER += 1
    elif upperChCount < 4:
        strength_UPPER += 2
    else:
        strength_UPPER += 3

    return strength_UPPER * RATE_UPPER


print(specialChCount)
print(numberChCount)
print(upperChCount)
strength = LEN(password) + NUMBER(password) + SPECIAL(password) + UPPER(password)
print(strength)


if upperChCount == 0 or specialChCount == 0 or numberChCount == 0:
    WEAK = "Uppercase" if upperChCount == 0 else "Special Character" if specialChCount == 0 else "Number"
    print(f"TOO WEAK. ENTER {WEAK} PLEASE")
else:
    strength = LEN(password) + NUMBER(password) + SPECIAL(password) + UPPER(password)
    if strength <= 1:
        print("WEAK")
    elif strength <= 2:
        print("MEHH")
    else:
        print("STRONG")