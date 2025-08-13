import java.util.Objects;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        while(true){
            Scanner input = new Scanner(System.in);
            String text = input.nextLine();


            if(Objects.equals(text, "q")){
                break;
            }

            Run run = new Run(text);
            run.run();
        }
    }
}
