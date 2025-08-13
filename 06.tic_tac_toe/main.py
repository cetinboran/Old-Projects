from Game import TicTacToe
from Player import Human,Computer

Tik = TicTacToe()

HumanPlayer = Human("o", Tik)
ComputerPlayer = Computer("x", Tik)



while Tik.notDraw:
    HumanPlayer.move()
    if HumanPlayer.winner == True: break
    if HumanPlayer.turn == "2": ComputerPlayer.turn = "1"
    ComputerPlayer.move()
    if ComputerPlayer.turn == "2": HumanPlayer.turn = "1"
    if ComputerPlayer.winner == True: break
    if Tik.notDraw == False: break
    

