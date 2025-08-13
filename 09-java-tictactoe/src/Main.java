public class Main {
    public static void main(String[] args) {
        Player one = new Player('x');
        Player two = new Player('o');

        Computer AI = new Computer(one.getSymbol());

        TicTacToe game = new TicTacToe(one, AI);
        game.gameLoop();

    }
}
