class Player:
    def __init__(self, username) -> None:
        
        self.username = username
        self.health = 3

    # def LevelUp(self):
    #     self.currentExp += (self.wins[0] * self.easyExp) + (self.wins[1] * self.mediumExp) + (self.wins[2] * self.hardExp)
        
    #     if self.currentExp >= self.maxExp:
    #         self.level += 1
    #         self.currentExp = 0
    #         self.maxExp *= self.ExpConstant
    #         input(f"LEVEL UP! {self.level}")
