import java.util.Random;

public class Computer {
    private char symbol;

    Computer(char playerSymbol){
        if(playerSymbol == 'X' || playerSymbol == 'x'){
            this.symbol = 'O';
        }
        else if(playerSymbol == 'O' || playerSymbol == 'o'){
            this.symbol = 'X';
        }
    }

    public int[] computerTurn(char[][] board){
        int[][] AllCoordinates = validCoordinates(board);
        Random rand = new Random();
        int choice = rand.nextInt(this.countValidCoordinates(board));

        return AllCoordinates[choice];
    }

    private int[][] validCoordinates(char[][] board){
        int[][] coordinates = new int[this.countValidCoordinates(board)][2];

        int k = 0;
        for(int i = 0; i < board.length; i++){
            for(int j = 0; j < board[0].length; j++){
                if(board[i][j] == ' '){
                    coordinates[k++] = new int[]{i, j};
                }
            }
        }

        return coordinates;
    }

    private int countValidCoordinates(char[][] board){
        int count = 0;
        for (char[] chars : board) {
            for (int j = 0; j < board[0].length; j++) {
                if (chars[j] == ' ') {
                    count++;
                }
            }
        }

        return count;
    }
    public char getSymbol(){
        return this.symbol;
    }
}
