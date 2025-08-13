import java.util.Scanner;

public class Player {
    private final char symbol;

    Player(char symbol){
        this.symbol = symbol;
    }

    public int[] myTurn(){
        byte row, col;

        do{
            Scanner input = new Scanner(System.in);

            System.out.print("Enter the row: ");
            row = input.nextByte();

            System.out.print("Enter the col: ");
            col = input.nextByte();
        } while(!(row <= 3 && col <= 3 && row > 0 && col > 0));

        return new int[] {row - 1,col - 1};
    }

    public char getSymbol(){
        return this.symbol;
    }
}
