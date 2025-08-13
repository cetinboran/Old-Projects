import argparse

ap = argparse.ArgumentParser()
ap.add_argument("-s", "--string", required=True, help="the text to encrypt/decrypt")
ap.add_argument("-n", "--number", required=True, help="number of characters to skip")
ap.add_argument("--encrypt", required=False, help="Encrypting", action='store_true')
ap.add_argument("--decrypt", required=False, help="Decrypting", action='store_true')
args = vars(ap.parse_args())

def encrypt(str, n=3):
    output = ""
    for ch in str:
        ascii_num = ord(ch)
        ascii_num += int(n)
        new_ch = chr(ascii_num)
        output += new_ch
    return output

def decrypt(str, n=3):
    output = ""
    for ch in str:
        ascii_num = ord(ch)
        ascii_num -= int(n)
        new_ch = chr(ascii_num)
        output += new_ch
    return output

if args["encrypt"] != True and args["decrypt"] != True:
    print("Enter one active attribute")
else:
    if args["encrypt"] != False:
        print(encrypt(args["string"], args["number"]))
    elif args["decrypt"] != False:
        print(decrypt(args["string"], args["number"]))
