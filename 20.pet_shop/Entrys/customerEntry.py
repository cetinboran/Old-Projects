from Entrys.login_loops import Login
from jsons.data import JSON

import os
class CustomerEntry(Login):
    def Loop(self, personList, animalsList, person):
        

        
        while True:
            os.system("cls")
            print(f"Login Successful. Hello {person.username}\n")

            print("1 - Show Animals")
            print("2 - Buy Animals")
            print("9 - Exit")
            command = input(">: ")

            if command == "9": break
            elif command == "1":
                os.system("cls")
                ShowAnimals(animalsList)
                input("")
            elif command == "2":
                os.system("cls")
                BuyAnimals(animalsList, person)
                JSON.UpdateAnimalsJson(JSON, animalsList)
                JSON.UpdatePersonsJson(JSON, personList)

def BuyAnimals(animalsList, person):
    ShowAnimals(animalsList)
    try:
        animalIdI = int(input("What Do You Want To Buy. Choose From ID: "))
        for animal in animalsList:
            if animal.animalId == animalIdI:
                if animal.cost <= person.balance:
                    print("Being bought")
                    person.balance -= animal.cost
                    if animal.type == "Dog":
                        person.dog += 1
                    elif animal.type == "Cat":
                        person.cat +=1
                        
                    #animalsList.remove(animal)
                else:
                    input("You dont have enought balance.")
                    break
                    
    except ValueError:
        input("Enter a Integer! ")

    # input("There is no such animal for this ID")



def ShowAnimals(animalsList):
    animal_info = ""
    for animal in animalsList:
        animal_info += f"Animal Id: {animal.animalId}, Name: {animal.name}, Type: {animal.type}, Rare: {animal.rare}, Cost: {animal.cost}\n"
    print(animal_info)