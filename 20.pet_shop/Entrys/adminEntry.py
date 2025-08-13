from Entrys.login_loops import Login
from jsons.data import JSON
from classes.person import *

import os

class AdminEntry(Login):
    def Loop(self, personList, customerList, workerList, person):
        os.system("cls")
        print(f"Login Successful. Hello Boss.\n")
        while True:
            print("1 - Hire Worker")
            print("2 - Fire Worker")
            print("3 - Show Workers")
            print("4 - Show Customers")
            print("9 - Exit")
            command = input(">: ")

            if command == "9": break
            elif command == "1":
                os.system("cls")
                HireWorker(personList, workerList)
                JSON.UpdatePersonsJson(JSON, personList)
            elif command == "2":
                os.system("cls")
                FireWorker(personList, workerList)
                JSON.UpdatePersonsJson(JSON, personList)
            elif command == "3":
                os.system("cls")
                ShowWorkers(workerList)
                input("")
            elif command == "4":
                os.system("cls")
                ShowCustomers(customerList)
                input("")

def HireWorker(personList, workerList):
    username = input("Enter the username: ")
    password = input("Enter the password: ")
    salary = input("Enter the salary: ")

    w = Worker(username, password, salary)
    personList.append(w)
    workerList.append(w)

def ShowWorkers(workersList):
    worker_info = ""
    for worker in workersList:
        worker_info += f"Worker Id: {worker.id}, Name: {worker.username}, Type: {worker.type}, Salary: {worker.balance}\n"
    print(worker_info)

def ShowCustomers(customerList):
    customer_info = ""
    for customer in customerList:
        customer_info += f"Customer Id: {customer.id}, Username: {customer.username}, Password: {customer.password}, Balance: {customer.balance}, Animals: {customer.animals} Type: {customer.type}, Salary: {customer.balance}\n"
    print(customer_info)


def FireWorker(personList, workerList):
    ShowWorkers(workerList)
    wId = int(input("Who You Gonna Fire. Choose From ID:"))
    for person in personList:
        if person.id == wId:
            print(f"{person.username} is fired")
            personList.remove(person)
            workerList.remove(person)