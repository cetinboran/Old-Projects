public class TicTacToe {
    private char[][] board;
    private final Player[] players = new Player[2];
    public Computer comp;
    private boolean gameOver;

    TicTacToe(){
        this.board = new char[][]{
                {' ', ' ', ' '},
                {' ', ' ', ' '},
                {' ', ' ', ' '},
        };

        this.gameOver = false;

        this.players[0] = new Player('X');
        this.players[1] = new Player('O');
    }
    TicTacToe(Player One, Player Two){
        this.board = new char[][]{
                {' ', ' ', ' '},
                {' ', ' ', ' '},
                {' ', ' ', ' '},
        };

        this.gameOver = false;

        if(One.getSymbol() == 'X' || One.getSymbol() == 'x'){
            this.players[0] = new Player(One.getSymbol());
            this.players[1] = new Player(Two.getSymbol());
        }
        else if(One.getSymbol() == 'O' || One.getSymbol() == 'o'){
            this.players[1] = new Player(One.getSymbol());
            this.players[0] = new Player(Two.getSymbol());
        }
        else{
            System.out.println("Invalid Symbol");
            System.exit(1);
        }
    }

    TicTacToe(Player player, Computer comp){
        this.board = new char[][]{
                {' ', ' ', ' '},
                {' ', ' ', ' '},
                {' ', ' ', ' '},
        };

        this.gameOver = false;

        if(player.getSymbol() == 'X' || player.getSymbol() == 'x'){
            this.players[0] = new Player('X');
            this.players[1] = null;
            this.comp = comp;
        }
        else if(player.getSymbol() == 'O' || player.getSymbol() == 'o'){
            this.players[1] = new Player('O');
            this.players[0] = null;
            this.comp = comp;
        }
    }

    public void gameLoop(){
        if(!checkSymbols()){
            System.out.println("Symbols cannot be the same!");
            return;
        }

        System.out.println("The Board");
        this.drawBoard();

        while(!this.gameOver){
            int[] coordinatesP1, coordinatesP2;

            if(players[0] != null){
                System.out.println("Player 1 Turn");
                do{
                    coordinatesP1 = players[0].myTurn();
                } while(this.validCoordinates(coordinatesP1));

                this.drawSymbols(players[0].getSymbol(), coordinatesP1);
                this.drawBoard();

                if(checkWin(players[0].getSymbol())){
                    System.out.printf("Player %s Won", players[0].getSymbol());
                    gameOver = true;
                    break;
                }
            }
            else{
                System.out.println("Computer Turn");

                coordinatesP1 = this.comp.computerTurn(this.board);
                this.drawSymbols(comp.getSymbol(), coordinatesP1);
                this.drawBoard();

                if(checkWin(comp.getSymbol())){
                    System.out.printf("Computer %s Won", comp.getSymbol());
                    gameOver = true;
                    break;
                }
            }
            if(players[1] != null){
                System.out.println("Player 2 Turn");
                do{
                    coordinatesP2 = players[1].myTurn();
                } while(this.validCoordinates(coordinatesP2));

                this.drawSymbols(players[1].getSymbol(), coordinatesP2);
                this.drawBoard();

                if(checkWin(players[1].getSymbol())){
                    System.out.printf("Player %s Won", players[1].getSymbol());
                    gameOver = true;
                    break;
                }

                if(checkDraw()){
                    System.out.println("Game is Draw!");
                    gameOver = true;
                    break;
                }
            }
            else{
                System.out.println("Computer Turn");

                coordinatesP2 = this.comp.computerTurn(this.board);
                this.drawSymbols(comp.getSymbol(), coordinatesP2);
                this.drawBoard();

                if(checkWin(comp.getSymbol())){
                    System.out.printf("Computer %s Won", comp.getSymbol());
                    gameOver = true;
                    break;
                }
            }
        }
    }

    private void drawSymbols(char symbol, int[] coordinates){
        int row = coordinates[0], col = coordinates[1];

        this.board[row][col] = symbol;
    }
    private boolean checkSymbols(){
        if(players[0] != null && players[1] != null){
            return this.players[0].getSymbol() != this.players[1].getSymbol();
        }
        return true;
    }

    private boolean validCoordinates(int[] coordinates){
        return this.board[coordinates[0]][coordinates[1]] != ' ';
    }

    public boolean checkDraw(){
        boolean draw = false;

        for (char[] chars : this.board) {
            for (int j = 0; j < this.board[0].length; j++) {
                if (chars[j] != ' ') {
                    draw = true;
                } else {
                    draw = false;
                    break;
                }
            }
            if (!draw)
                break;
        }

        return draw;
    }
    private boolean checkWin(char Symbol){
        return checkWinDiag(Symbol) || checkWinStraightBelow(Symbol);
    }
    private boolean checkWinStraightBelow(char Symbol){
        boolean win = false;

        for (char[] chars : this.board) {
            for (int i = 0; i < this.board[0].length; i++) {
                if (chars[i] == Symbol) {
                    win = true;
                } else {
                    win = false;
                    break;
                }
            }
            if (win)
                return win;
        }

        for(int k = 0; k < this.board.length; k++) {
            for (int i = 0; i < this.board[0].length; i++) {
                if (this.board[i][k] == Symbol) {
                    win = true;
                } else {
                    win = false;
                    break;
                }
            }
            if(win)
                return win;
        }

        return false;
    }
    private boolean checkWinDiag(char Symbol){
        boolean win = false;

        // Left to Right
        for(int i = 0; i < this.board.length; i++){
            if(this.board[i][i] == Symbol){
                win = true;
            }
            else{
                win = false;
                break;
            }
        }

        if(win)
            return win;

        // Right to left
        int k = 2;
        for (char[] chars : this.board) {
            if (chars[k] == Symbol) {
                win = true;
            } else {
                win = false;
                break;
            }
            k--;
        }

        if(win)
            return win;

        return false;
    }

    private void drawBoard(){
        int count = 0;
        for(char[] row: this.board){
            for(char x: row){
                System.out.print(" " + x + " ");
                if(count % 3 == 0 || count % 3 == 1)
                    System.out.print("| ");
                count++;
            }
            System.out.println();
        }
        System.out.println();
    }
}
