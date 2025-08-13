from Game import TicTacToe
import random as r
import os

class Player:
    def __init__(self,letter) -> None:
        self.letter = letter

    def win(self, game, letter):
        board = TicTacToe.return_board(game)
        sag = ""
        sol = ""
        x = 2
        for row in range(0, 3):
            yatay = ""
            dikey = ""
            for col in range(0, 3):
                yatay += board[row][col]
                dikey += board[col][row]
            

            sag += board[row][row]
            sol += board[row][x]
            x -= 1

            if yatay.count(self.letter) == 3 or dikey.count(self.letter) == 3:
                return True
            
        if sag.count(self.letter) == 3 or sol.count(self.letter) == 3:
            return True


    def find_available_moves(self):
        availableMoves = []
        board = TicTacToe.return_board(self.game)

        for row in range(3):
            for col in range(3):
                if board[row][col] == " ":
                    availableMoves.append((row, col)) # Kordinatları Tuple olarak attın
        return availableMoves

    def convertedMove(self,move):
        convertedMove = None
        if move == 1: convertedMove = (0, 0)
        elif move == 2: convertedMove = (0, 1)
        elif move == 3: convertedMove = (0, 2)
        elif move == 4: convertedMove = (1, 0)
        elif move == 5: convertedMove = (1, 1)
        elif move == 6: convertedMove = (1, 2)
        elif move == 7: convertedMove = (2, 0)
        elif move == 8: convertedMove = (2, 1)
        elif move == 9: convertedMove = (2, 2)
        if convertedMove != None:
            return convertedMove

class Human(Player):

    def __init__(self, letter, game) -> None:
        super().__init__(letter)
        self.game = game
        self.turn = ""
        self.winner = False

        if letter == "x": self.turn = "1"
        else: self.turn = "2"

    def win(self, game, letter):
        return super().win(game, letter)

    def find_available_moves(self):
        return super().find_available_moves()

    def convertedMove(self, move):
        return super().convertedMove(move)

    def move(self):
        board = TicTacToe.return_board(self.game)
       
        notDraw = True
        notValid = True
        while notValid and notDraw:
            availableMoves = self.find_available_moves()
            

            if self.turn == "2":
                break

            try:
                if availableMoves == []:
                    notDraw = False
                    TicTacToe.set_draw(self.game, notDraw)
                    os.system("cls")
                    TicTacToe.draw_board(self.game)
                    input("This game is a draw")
                    break
                else:
                    os.system("cls")

                    print("*******************************")
                    TicTacToe.draw_board(self.game)
                    print("*******************************")

                    # isOver = TicTacToe.return_winner(self.game)
                    # print(isOver)
                    
                    move = int(input(f"Select position of your {self.letter} (1 - 9): "))
                    convertedMove = self.convertedMove(move)

                    for moves in availableMoves:
                        if convertedMove == moves:
                            self.turn = "2"
                            board[convertedMove[0]][convertedMove[1]] = self.letter
                            notValid = False
                            break
                    
            except:
                input("Please, enter valid move") # try except döngüsünden çıkamadığımdan break attım
            
        winner = self.win(self.game, self.letter)
        if winner:
            self.winner = True
            os.system("cls")
            TicTacToe.draw_board(self.game)
            print(f"{self.letter} is winner!")

        # Oyunun Board'ını Game Classı için Setledim.
        TicTacToe.set_board(self.game, board)


class Computer(Player):
    def __init__(self, letter, game) -> None:
        super().__init__(letter)
        self.game = game
        self.turn = ""
        self.winner = False

        if letter == "x": self.turn = "1"
        else: self.turn = "2"
    
    def find_available_moves(self):
        return super().find_available_moves()
    
    def win(self, game, letter):
        return super().win(game, letter)

    def convertedMove(self, move):
        return super().convertedMove(move)

    def move(self):
        board = TicTacToe.return_board(self.game)

        notDraw = True
        notValid = True
        while notValid and notDraw:
            availableMoves = self.find_available_moves()
           
            if self.turn == "2":
                break
            try:
                if availableMoves == []:
                    notDraw = False
                    TicTacToe.set_draw(self.game, notDraw)
                    os.system("cls")
                    TicTacToe.draw_board(self.game)
                    input("This game is a draw")
                    break
                else:
                    move = r.randint(1,9)
                    convertedMove = self.convertedMove(move)

                    for moves in availableMoves:
                        if convertedMove == moves:
                            self.turn = "2"
                            board[convertedMove[0]][convertedMove[1]] = self.letter
                            notValid = False
                            break
            except:
                input("Please, enter valid move") # try except döngüsünden çıkamadığımdan break attım


        winner = self.win(self.game, self.letter)

        if winner:
            self.winner = True
            os.system("cls")
            TicTacToe.draw_board(self.game)
            print(f"{self.letter} is winner!")

        TicTacToe.set_board(self.game, board)