import os
import json

class Customer:
    # Global CustomerId oluşturdum ve bi altında atama yaptım
    global customerId 
    customerId = 1
    
    def __init__(self, Name, Age) -> None:
        # global customerId'nin referansını buraya çektim ve id işlemlerini yaptım.
        global customerId
        
        self.id = customerId
        customerId += 1
        
        self.name = Name
        self.age = Age

customers = []

def CreateCustomer(cust=list):
    name = input("Enter Customer Name: ")
    age = input("Enter Customer Age: ")

    if(name != "" and age != ""):
        cust.append( Customer(name, age))
    else:
        input("Please enter some information!")

    os.system("cls")

def UpdateCustomer(cust=list):
    for customer in cust:
        print("-----------------------------")
        print(f"Customer Id: {customer.id} ")
        print(f"Customer Name: {customer.name} ")
        print(f"Customer Age: {customer.age} ")
        print("-----------------------------")
    
    customerId = int(input("Enter the id of the customer to be updated: "))
    name = input("Enter Customer Name: ")
    age = input("Enter Customer Age: ")

    cust[customerId - 1].name = name
    cust[customerId - 1].age = age
    os.system("cls")
    

def ReadCustomers(cust=list):
    for customer in cust:
        print("-----------------------------")
        print(f"Customer Id: {customer.id} ")
        print(f"Customer Name: {customer.name} ")
        print(f"Customer Age: {customer.age} ")
        print("-----------------------------")

    input("Press to continue...")
    os.system("cls")

def LoadCustomers(cust=list):
    with open("customer.json") as f:
        data = json.load(f)

    # print(data)
    for customer in data:
        txt = ""

        for info in customer:
            txt += str(customer[info])
        cust.append( Customer(txt[1], txt[2]))



def SaveCustomers(cust=list):
    data = []
    for customer in cust:
        data_dict = {"id": customer.id, "name": customer.name, "age": customer.age}
        data.append(data_dict)

        with open("customer.json","w") as json_file:
            json.dump(data, json_file, indent=4, sort_keys=False)


def Start():
    LoadCustomers(customers)
    while True:
        print("1 - Create")
        print("2 - Update")
        print("3 - Read")
        print("4 - Save")
        print("6 - Quit")
        command = input(">: ")
        
        os.system("cls")

        if(command == "1"):
            CreateCustomer(customers)
        elif(command == "2"):
            UpdateCustomer(customers)
        elif(command == "3"):
            ReadCustomers(customers)
        elif(command == "4"):
            SaveCustomers(customers)
        elif(command == "6"):
            break
        else:
            print("Invalid Command")
        

Start()