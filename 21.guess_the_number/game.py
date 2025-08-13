import random as r
import os, json
from types import NoneType
from player import Player

player_list = []

def player_init():
    with open("players.json", "r") as json_file:
        players = json.load(json_file)
        
    for player in players:
        pl = Player(player["username"])
        player_list.append(pl)

def update_JSON():
    data = []
    for player in player_list:
        player_dict = {"username": player.username, "health": player.health }
        data.append(player_dict)

        with open("players.json", "w") as json_file:
            json.dump(data,json_file, sort_keys=True, indent=4)
    
def LoginSection():
    while True:
        os.system("cls")
        print("1 - Login")       
        print("2 - Register")     
        print("3 - exit")  
        command = input(">: ")

        if command == "1":
            return Login()
        elif command == "2":
            Register()
        elif command == "3":
            return None

def Login():
    username = input("Enter your username: ")
    for player in player_list:
        if player.username == username:
            return player

    input("Wrong username, please try again.")

def Register():
    username = input("Enter your username: ")
    gamer = Player(username)
    for player in player_list:
        if player.username == username:
            input("Name already taken. Please choose different username")
        else:
            player_list.append(gamer)
            input("Your account has been successfully created. Have fun! ")
            break
    if len(player_list) == 0:
        player_list.append(gamer)
        input("Your account has been successfully created. Have fun! ")
    update_JSON()

def ChooseDifficulty():
    while True:
        os.system("cls")
        print("1 - Easy")       
        print("2 - Medium")     
        print("3 - Hard")      
        print("4 - Exit")      
        command = input(">: ")
        
        if command.strip() == "1": 
            return 5
        elif command.strip() == "2":
            return 10
        elif command.strip() == "3": 
            return 15
        elif command.strip() == "4": 
            break

def StartGame(maxValue, player):
    computerRandomNumber = r.randint(1, maxValue)
    win = False

    while player.health > 0:
        try:

            playerGuess = int(input("Enter your guess: "))
            if playerGuess > computerRandomNumber: 
                os.system("cls")
                print("Too High!")
                player.health -= 1
                print(f"Remaining Health: {player.health}")
            elif playerGuess < computerRandomNumber: 
                os.system("cls")
                print("Too Low")
                player.health -= 1
                print(f"Remaining Health: {player.health}")
            else:
                input("You Win!")
                break

           
                
                        
            if player.health == 0: 
                input("You Lost!")
        except:
            print("The guess has to be a Number")
    player.health = 3


def GameLoop():
    player_init()

    Login = ""
    while True:
        os.system("cls")
        if type(Login) == str:
            print("1 - Start")
            print("2 - About")
            print("3 - Exit")
            command = input(">: ")
        
        if command.strip() == "1": 
            if type(Login) == str:
                Login = LoginSection()

            if type(Login) == NoneType: break

            maxValue = ChooseDifficulty()
            if type(maxValue) == NoneType: break

            StartGame(maxValue, Login)
        elif command.strip() == "2": pass
        elif command.strip() == "3": break


GameLoop()




