class Animal:
    global animalId
    animalId = 1
    def __init__(self) -> None:
        global animalId
        self.animalId = animalId
        animalId += 1

        
class Dog(Animal):
    global dogId
    dogId = 1
    def __init__(self, name, rare, cost) -> None:
        super().__init__()
        global dogId
        self.id = dogId
        dogId += 1
        
        self.name = name
        self.type = "Dog"
        self.rare = rare
        self.cost = cost
        
class Cat(Animal):
    global catId
    catId = 1
    def __init__(self, name, rare, cost) -> None:
        super().__init__()
        global catId
        self.id = catId
        catId += 1
        
        self.name = name
        self.type = "Cat"
        self.rare = rare
        self.cost = cost