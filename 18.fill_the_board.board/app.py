import random as r
import os

try:
    BOARDLENGTH = int(input("Enter the Size of The Board: "))

    if BOARDLENGTH == 1:
        print("Board Length Cannot Be 1")
        quit()
except:
    print("Please Enter Integet")
    quit()


board = [[" " for i in range(0,BOARDLENGTH)] for i in range(0,BOARDLENGTH)]
emptySquares = []

global passedCount
passedCount = 1

try :
    StartPointX = int(input(f"(0 - {BOARDLENGTH - 1}) Choose X Position: "))
    StartPointY = int(input(f"(0 - {BOARDLENGTH - 1}) Choose Y Position: "))

    if StartPointX > BOARDLENGTH or StartPointY > BOARDLENGTH:
        print("Board too Small For That Input.")
        quit()
except ValueError:
    print("Please Enter Integer")
    quit()

START = [StartPointX, StartPointY]

def DrawBoard():
    for row in board:
        output = " "
        for col in row:
            output += col
        print(output)

def Passed(emptySquares, board , exPosition):
    global passedCount

    for row in range(0, len(board), 1):
        for col in range(0, len(board), 1):
            if exPosition[0] == row and exPosition[1] == col:
                if BOARDLENGTH >= 10:
                    if passedCount < 10:
                        board[row][col] = "   " + f"000{passedCount}" + "   "
                    elif passedCount >= 100:
                        board[row][col] = "   " + f"0{passedCount}" + "   "
                    else:
                        board[row][col] = "   " + f"00{passedCount}" + "   "
                else:
                    if passedCount < 10:
                        board[row][col] = "   " + f"0{passedCount}" + "   "
                    else:
                        board[row][col] = "   " + str(passedCount) + "   "

    passedCount += 1
    if len(emptySquares) != 1:
        emptySquares.remove([exPosition[0], exPosition[1]])

def FindEmptySquare(emptySquares, board):
    for row in range(0, len(board), 1):
        for col in range(0, len(board), 1):
            emptySquares.append([row,col])
            
            

def Start(START):
    while True:
        os.system("cls")
        DrawBoard()
        choose = r.choice(emptySquares)

        exPosition = START
        if choose in emptySquares and choose != exPosition:
            Passed(emptySquares, board, exPosition)
            START[0], START[1] = choose[0], choose[1]
        elif len(emptySquares) == 1:
            Passed(emptySquares, board, exPosition)
            
            os.system("cls")
            DrawBoard()
            break

        os.system("cls")
        DrawBoard()
        if len(emptySquares) == 0: break
    
        
FindEmptySquare(emptySquares, board)
Start(START)

