from Entrys.workerEntry import WorkerEntry
from Entrys.customerEntry import CustomerEntry
from Entrys.adminEntry import AdminEntry

def Login(personList):
    username = input("Enter your username: ")
    password = input("Enter your password: ")
    for person in personList:
        if person.type == "Worker" and person.username == username and person.password == password:
            return [WorkerEntry, person]
        elif person.type == "Customer" and person.username == username and person.password == password:
            return [CustomerEntry, person]
        elif person.type == "Admin" and person.username == username and person.password == password:
            return [AdminEntry, person]

    print("username or password is incorrect")
    return [False,False]