class Person:
    global pId
    pId = 1

    def __init__(self, username, password, balance) -> None:
        global pId
        self.id = pId
        pId += 1

        self.username = username
        self.password = password
        self.balance = balance

 

class Worker(Person):
    def __init__(self, username, password, balance,) -> None:
        super().__init__( username, password, balance)
        self.type = "Worker"

class Customer(Person):
    def __init__(self, username, password, balance, dog, cat) -> None:
        super().__init__(username, password, balance)
        self.type = "Customer"

        self.dog = dog
        self.cat = cat
        self.animals = [
            self.dog,
            self.cat
        ]

class Admin(Person):
    def __init__(self, username, password, balance) -> None:
        super().__init__(username, password, balance)
        self.type = "Admin"
        