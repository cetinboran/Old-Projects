import subprocess
import os

current_path = os.getcwd()

file_input = input("File: ")
copy_type = input("Kopya Yazilisi: ")

file_name = os.path.join(current_path, file_input)

# /b name'leri alÃ„Â±yor.
command = subprocess.run(f"dir {file_name} /b", capture_output=True, shell=True)
output = command.stdout.decode().strip().replace("\r", "").split("\n")


for file in output:
    if copy_type in file:
        new_path = os.path.join(file_name, file)
        #print(f"del '{new_path}'")
        subprocess.run(f'del "{new_path}"', shell=True)