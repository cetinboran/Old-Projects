package Commands;

import java.util.Arrays;
import java.util.Objects;

public class Wc extends Command{
    String[] validOptions;
    boolean[] enteredOptions;
    public Wc(String name, String args, String options) {
        super(name, args, options);

        this.validOptions = new String[] {"-l", "-w", "-c"};
        this.enteredOptions = new boolean[] {false, false, false};

        setStdI(false);
    }

    @Override
    public void run() {
        handleOptions();
        start();
    }

    private int[] handleWc(){
        int[] output = new int[3];
        int lineCount = 1;
        int wordCount = 0;
        int charCount = 0;

        boolean isWord = false;
        if(getStdI()){

            if(getInput() == null){
                setInput(" ");
            }

            for(int i = 0; i < getInput().length(); i++){
                char c = getInput().charAt(i);
                if((c != ' ') && !isWord) {
                    isWord = true;
                    wordCount++;
                }
                else if(c == ' ' || c == '\n'){
                    isWord = false;
                }
                if(c == '\n'){
                    lineCount++;
                }

                charCount++;
            }

            output[0] = lineCount;
            output[1] = wordCount;
            output[2] = charCount;
        }

        return output;
    }

    private void start(){
        String out = calculateOutput(handleWc());

        // STDO 0 olduğu için çıktı ekrana gider
        // İnput yok ise ekrana yazmasın dedik.

        if(!getStdO() && getStdI()){
            System.out.println(out);
        }
        else{
            // STD1 ise pipe'ın outputu alıp diğer komuta atması için değişkene atıyoruz.
            setOutput(out);
        }
    }

    private String calculateOutput(int[] output){
        int lineCount = output[0];
        int wordCount = output[1];
        int charCount = output[2];

        String out = "";

        if(this.enteredOptions[0])
            out += lineCount + " ";

        if(this.enteredOptions[1])
            out += wordCount + " ";

        if(this.enteredOptions[2])
            out += charCount + " ";

        if(!this.enteredOptions[0] && !this.enteredOptions[1] && !this.enteredOptions[2]){
            out = lineCount + " " + wordCount + " " + charCount;
        }

        return out;
    }

    private void handleOptions(){
        if(this.options == null)
            return;

        for(int i = 0; i < this.validOptions.length; i++){
            for (String option : this.options) {
                if (Objects.equals(validOptions[i], option)) {
                    this.enteredOptions[i] = true;
                    break;
                }
            }
        }
    }
}
