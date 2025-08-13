package Commands;

import java.util.Arrays;
import java.util.Scanner;

public class Echo extends Command {
    public Echo(String name, String args, String options) {
        super(name, args, options);
    }

    public void run(){
        emptyOption();
    }

    private void emptyOption(){
        // Eğer input var ve pipe yok ise ekrana yaz.
        String output = "";

        if(getArgs() == null){
            setArgs(new String[] {""});
        }

        if(getStdI()){
            if(!getStdO()){
                // Ekrana yazıyoruz SDTO 0
                for(String input: getArgs()){
                    output += input + " ";
                }

                System.out.println(output);
            }
            else{
                for(String input: getArgs()){
                    output += input + " ";
                }
                // Pipe İçin kaydediyoruz STDO 1
                setOutput(output);
            }
        }
        else{
            // Eğer input yok ise
            if(!getStdO()){
                // Ekrana boşuk yazar çünkü STDO 0
                System.out.println();
            }
            else {
                // Boşluk kaydeder çünkü STD0 1
                setOutput("");
            }
        }
    }

    private void stdIFalse(){
        System.out.println();
    }

}
