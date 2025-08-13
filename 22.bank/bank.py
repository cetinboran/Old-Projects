import json
import os
from accounts import Account
from customer import Customer

class Bank():
    
    def __init__(self, Name) -> None:
        self.name = Name

        self.accounts_list = []             # All Accounts
        self.customers_accounts_list = []   # Just Customer's Accounts
        self.target_accounts_list = []      # Target Accounts 
        self.customers_list = []            # Customers

        self.init_JSON(self.accounts_list, self.customers_list)


    def init_JSON(self, accounts_list, customers_list):
        with open("Accounts.json", "r") as json_file:
            accounts = json.load(json_file)
            for account in accounts:
                    self.accounts_list.append(Account(account["customerId"], account["balance"], account["accountName"]))

        with open("Users.json", "r") as json_file:
            customers = json.load(json_file)
            for customer in customers:
                    self.customers_list.append(Customer(customer["username"], customer["password"]))

    def update_customers_JSON(self):
        customers_list = []

        for customer in self.customers_list: 
            customer_dict = {"id": customer.id, "username": customer.username, "password": customer.password}
            customers_list.append(customer_dict)
            with open("Users.json", "w") as json_file:
                json.dump(customers_list, json_file, indent=4, sort_keys=False)

    def update_accounts_JSON(self):
        accounts_list = []

        for account in self.accounts_list:
            account_dict = {"customerId": account.customer_id, "id": account.id, "balance": account.balance, "accountName": account.accountName}
            accounts_list.append(account_dict)
            with open("Accounts.json", "w") as json_file:
                json.dump(accounts_list, json_file, indent=4, sort_keys=False)

    def Authentication(self):
        os.system("cls")
        username = input("Enter your username: ")
        password = input("Enter your password: ")

        for customer in self.customers_list:
            if customer.username == username.lower() and customer.password == password.lower():
                return customer
            else:
                continue

    def Register(self):
        username = input("Enter your username: ")
        password = input("Enter your password: ")

        if len(username) > 2 and len(password) > 2:
            cust = Customer(username, password)
            self.customers_list.append(cust)
            
            self.update_customers_JSON() 

            accountName = input("Enter your account name: ")
            if len(accountName) > 2:
                customer_account = Account(cust.id, 0, accountName)
                self.accounts_list.append(customer_account)

                self.update_accounts_JSON()

                input("Registration Successful")
            else: input("Account name must be 3 characters long.")
        else: input("Username and password must be 3 characters long.")

    def findCustomerAccounts(self, customer):
        for account in self.accounts_list:
            if account.customer_id == customer.id:
                self.customers_accounts_list.append(account)
            else:
                self.target_accounts_list.append(account)

    def select_target_account(self):
        os.system("cls")

        for account in self.target_accounts_list:
            print(f"Id: {account.id} | Balance: {account.balance} | Accout Name: {account.accountName}")

        try:
            id = int(input("Choose from your accounts\nEnter your target account id: "))
            for account in self.target_accounts_list:
                if account.id == id:
                    return account
                    break
        except:
            input("Invalid Entry")

    def select_customer_account(self):
        os.system("cls")

        for account in self.customers_accounts_list:
            print(f"Id: {account.id} | Balance: {account.balance} | Accout Name: {account.accountName}")
        print("*************************************************")

        try:
            id = int(input("Choose from your accounts\nEnter your account id: "))
            for account in self.customers_accounts_list:
                if account.id == id:
                    return account
                    break
            input("Invalid ID")
        except:
            input("Invalid Entry")

    def withdraw_money(self):
        customer_account = self.select_customer_account()

        if customer_account != None:
            try:
                money = int(input("How much are you going to withdraw: "))
                if money > 0:
                    if customer_account.balance >= money:
                        customer_account.balance -= money

                        self.update_accounts_JSON()
                        input("successfully withdrawn money!")
                        input(f"Remaining balance: {customer_account.balance}")
                    else:
                        input("You don't have enough money in your account")
                else: input("You cant withdraw negative amount!")
            except:
                input("Invalid Entry")

    def show_accounts(self):
        for account in self.customers_accounts_list:
            print(f"Id: {account.id} | Balance: {account.balance} | Accout Name: {account.accountName}")
        input("*******************************************")

    def deposit_money(self):
        customer_account = self.select_customer_account()
        
        if customer_account != None:
            try:
                money = int(input("How much are you going to deposit: "))
                if money > 0:
                    customer_account.balance += money

                    self.update_accounts_JSON()
                    input("successfully deposited!")
                    input(f"Balance: {customer_account.balance}")
                else: input("You cant deposit negative amount!")
            except:
                input("Invalid Entry")
        else: input("You have to select a valid account")
    
    def transfer_money(self):
        customer_account = self.select_customer_account()
        if customer_account != None:
            target_account = self.select_target_account()
            if target_account != None:
                try:
                    money = int(input("Enter the money: "))
                    if money > 0:
                        if customer_account.balance >= money:
                            print("Transfer is initiating...")
                            customer_account.balance -= money
                            target_account.balance += money
                            print("Transfer completed")
                            input(f"Remaining Balance: {customer_account.balance} ")

                            self.update_accounts_JSON()
                        else: input("You dont have enought money")
                    else: input("You cant transfer negative amount")
                except:
                    print("Invalid Entry")
            else: input("You have to enter a valid target id")
        else: input("You have to enter a valid target id")

    def create_new_account(self, customer):
        accountName = input("Enter your account name: ")
        newAccount = Account(customer.id, 0, accountName)
        self.accounts_list.append(newAccount)
        self.update_accounts_JSON()
        self.customers_accounts_list.append(newAccount)


    def Start(self):
        customer = None

        Login = True
        while Login:
            os.system("cls")
            print("1 - Login")
            print("2 - Register")
            print("9 - Quit")
            command = input(">: ")
            if(command == "1"): 
                customer = self.Authentication()
                if customer == None: input("Username or Password is Incorrect.")
                else: break
            elif(command == "2"):
                self.Register()
            elif(command == "9"): 
                Login = False
        
        if customer != None: self.findCustomerAccounts(customer)

        while customer != None:
            os.system("cls")
            print("1 - Show Accounts")
            print("2 - Withdraw Money")
            print("3 - Deposit Money")
            print("4 - Transfer Money")
            print("5 - Create New Account")
            print("9 - Quit")
            command = input(">: ")

            if(command == "9"): customer = None
            elif(command == "1"): 
                self.show_accounts()
            elif(command == "2"):
                self.withdraw_money()
            elif(command == "3"): 
                self.deposit_money()
            elif(command == "4"): 
                self.transfer_money()
            elif(command == "5"):
                self.create_new_account(customer)