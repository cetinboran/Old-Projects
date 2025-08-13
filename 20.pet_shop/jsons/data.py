from classes.animals import *
from classes.person import *
import json

class JSON:
    def init_animals(self):
        animals = JSON.GetJson("animals")

        animalsList = []
        dogList = []
        catList = []

        for animal in animals:
            if animal["type"] == "Dog":
                dog = Dog(animal["name"],animal["rare"],animal["cost"])
                dogList.append(dog)
                animalsList.append(dog)
            elif animal["type"] == "Cat":
                cat = Cat(animal["name"],animal["rare"],animal["cost"])
                catList.append(cat)
                animalsList.append(cat)
        return [animalsList, dogList, catList]

    def UpdateAnimalsJson(self, animalsList):
        data = []
        for animal in animalsList:
            animals_dict = {"animalId": animal.animalId , "id": animal.id, "name": animal.name, "type": animal.type, "rare": animal.rare , "cost": animal.cost }
            data.append(animals_dict)

        with open("jsons/animals.json", "w") as file:
            json.dump(data, file, indent=4)

    def UpdatePersonsJson(self, personsList):
        data = []
        for person in personsList:
            if person.type == "Worker":
                worker_dict = {"id": person.id, "username": person.username, "password": person.password, "salary": person.balance, "type": person.type}
                data.append(worker_dict)
            elif person.type == "Customer":
                customer_dict = {"id": person.id, "username": person.username, "password": person.password, "balance": person.balance, "type": person.type, "animals": { "dog": person.dog, "cat": person.cat}}
                data.append(customer_dict)
            elif person.type == "Admin":
                admin_dict = {"id": person.id, "username": person.username, "password": person.password, "type": person.type}
                data.append(admin_dict)


        with open("jsons/personsInfo.json", "w") as file:
            json.dump(data, file, indent=4)


    def init_persons(self):
        personsInfo = JSON.GetJson("personsInfo")

        admin = []
        worker = []
        customer = [] 
        personsList = []

        for person in personsInfo:
            if person["type"] == "Worker":
                w = Worker(person["username"], person["password"], person["salary"])
                worker.append(w)
                personsList.append(w)
            elif person["type"] == "Customer":
                c = Customer(person["username"],person["password"], person["balance"], person["animals"]["dog"],person["animals"]["cat"])
                customer.append(c)
                personsList.append(c)
            elif person["type"] == "Admin":
                d = Admin(person["username"], person["password"], 0)
                admin.append(d)
                personsList.append(d)
                
        return [admin, worker, customer, personsList]


    @staticmethod
    def GetJson(which):
        try:
            with open(f"jsons/{which}.json", "r") as file:
                animals = json.load(file)
            return animals
        except:
            print("There is no such file")