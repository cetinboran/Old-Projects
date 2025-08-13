class TicTacToe:
    def __init__(self) -> None:
        # 2D List Oluşturduk Tek Satırda
        self.board = [[" " for j in range(3)] for i in range(3)]

        self.notDraw = True

    def draw_board(self):
        for row in range(0, 3):
            output = ""
            for col in range(0, 3):
                output += " | " + self.board[row][col]+ " | "
            print(output + "\n")

    def set_draw(self, draw):
        self.notDraw = draw

    def return_board(self):
        board = self.board
        return board
    
    def set_board(self, board):
        self.board = board

