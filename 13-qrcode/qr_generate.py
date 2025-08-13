#!/usr/bin/python3

import qrcode
import optparse

EXTENSIONS = [".jpg", ".png", ".jpeg", ".gif"]

def check_version(version):
    if version > 40: 
        print("Version Cannot Bigger Than 40")
        exit(0)

def check_extention(name):
    if "." not in name[-4::]: 
        name += ".jpg" 
        print("You Dont Have Extention!")
        print("Default Extention is JPG!")

        return name

    find_dot = list(name).index(".")
    if name[find_dot::] in EXTENSIONS:
        return name
    else: print("Not Allowed Extention")

    exit(0)


def make_qr(size, string, name):
    size = int(size)
    check_version(size)
    name = check_extention(name)

    qr = qrcode.QRCode(version=size, box_size=10, border=(size)//2)
    qr.add_data(string)
    qr.make(fit=True)
    
    image = qr.make_image()
    image.save(name)


def main():
    parser = optparse.OptionParser("Usage of program: " + " -v <Size Of The QR (1 - 5)> -s <String Or Url> -n <File Name To Save>")
    parser.add_option("-v", dest="size", default="2", type='string', help="Set the size of QR code.")
    parser.add_option("-s", dest="string", type='string', help="Set the string for te QR code.")
    parser.add_option("-n", dest="name_qr", type='string', help="Set the name of the file you want to save with the extension")
    (options, args) = parser.parse_args()
    size = options.size
    string = options.string
    name_qr = options.name_qr


    if size == None or string == None and name_qr == None: 
        print(parser.usage)
        exit(0)

    make_qr(size, string, name_qr)

main()
