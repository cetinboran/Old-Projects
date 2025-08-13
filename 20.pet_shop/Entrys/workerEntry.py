from Entrys.login_loops import Login
from classes.animals import Dog, Cat
from jsons.data import JSON

import os
class WorkerEntry(Login):
    def Loop(self, dogList, catList, animalsList, person):
        

        
        while True:
            os.system("cls")
            print(f"Login Successful. Hello {person.username}\n")

            print("1 - Add Animal")
            print("2 - Update Animal")
            print("3 - Show Animals")
            print("9 - Exit")
            command = input(">: ")

            if command == "9": break
            elif command == "1":
                os.system("cls")
                AddAnimal(animalsList, dogList, catList)
                JSON.UpdateAnimalsJson(JSON,animalsList)
            elif command == "2":
                os.system("cls")
                UpdateAnimals(animalsList)
                JSON.UpdateAnimalsJson(JSON, animalsList)
            elif command == "3":
                os.system("cls")
                ShowAnimals(animalsList)
                input("")
            else:
                input("There is no such a menu")


def AddAnimal(animalList, dogList, catList):
    name = input("Enter the name of the animal: ")
    rare = input("Enter the rare of the animal: ")
    cost = input("Enter the cost of the animal: ")
    type = input("Choose Type (Dog / Cat): ")
    if len(name) > 2 and len(rare) > 2:
        if type.capitalize() == "Dog":
            d = Dog(name, rare, cost)
            animalList.append(d)
            dogList.append(d)
        elif type.capitalize() == "Cat":
            c = Cat(name, rare, cost)
            animalList.append(c)
            catList.append(c)
        else:
            input("There is no such type")
    else:
        input("Please at least enter 2 character")

def UpdateAnimals(animalsList):
    ShowAnimals(animalsList)
    try:
        animalId = int(input("Choose From ID: "))

        for index,animal in enumerate(animalsList):
            if animal.animalId == animalId:

                name = input("Enter the new name: ")
                rare = input("Enter the new rare: ")
                cost = input("Enter the new cost: ")
                if len(name) > 2 and len(rare) > 2:
                    animalsList[index].name = name
                    animalsList[index].rare = rare
                    animalsList[index].cost = cost
                else:
                    input("Please enter at least 2 character")
    except:
        input("Please Enter Integer")


def ShowAnimals(animalsList):
    animal_info = ""
    for animal in animalsList:
        animal_info += f"Animal Id: {animal.animalId}, Name: {animal.name}, Type: {animal.type}, Rare: {animal.rare}, Cost: {animal.cost}\n"
    print(animal_info)