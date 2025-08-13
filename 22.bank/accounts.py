class Account:
    global Id
    Id = 1

    def __init__(self, customerId, Balance, AccountName) -> None:
        global Id
        self.id = Id
        Id += 1
        self.customer_id = customerId
        self.balance = Balance
        self.accountName = AccountName
        