class Customer:
    global Id
    Id = 1

    def __init__(self, Username, Password) -> None:
        global Id
        self.id = Id
        Id += 1

        self.username = Username
        self.password = Password