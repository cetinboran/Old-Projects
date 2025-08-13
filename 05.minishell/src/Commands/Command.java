package Commands;

import java.util.Arrays;
import java.util.Objects;

public abstract class Command {
    private final int commandId;
    private final String name;
    private String[] args;
    public String[] options;

    private String input, output;
    private boolean stdI, stdO;

    public Command(String name, String args, String options){
        this.commandId = 1;
        this.name = name;
        setArgs(args.split(" "));
        setOptions(options.split(" "));

        // STDI 0 ise input yok 1 ise var
        // STDO 0 ise çıktıyı ekrana versin 1 ise pipe'a

        setStdI(!(getArgs() == null));
    }

    public abstract void run();

    public String getInput() {
        return input;
    }

    public void setInput(String input) {
        this.input = input;
    }

    public String getOutput() {
        return output;
    }

    public void setOutput(String output) {
        this.output = output;
    }

    public boolean getStdI() {
        return stdI;
    }

    public void setStdI(boolean stdI) {
        this.stdI = stdI;
    }

    public boolean getStdO() {
        return stdO;
    }

    public void setStdO(boolean stdO) {
        this.stdO = stdO;
    }

    public String[] getArgs() {

        return args;
    }

    public void setArgs(String[] args) {
        if(args.length == 1 && args[0].length() == 0){
            this.args = new String[] {" "};
        }
        else
            this.args = args;
    }

    public String[] getOptions() {
        return options;
    }

    public void setOptions(String[] options) {
        if(options.length == 1){
            if(Objects.equals(options[0], "")){
                this.options = null;
                return;
            }
        }


        String validOptions = "";
        for (String option : options) {
            if (!Objects.equals(option, "-")) {
                validOptions += option + " ";
            }
        }

        this.options = validOptions.trim().split(" ");
    }

    public int getId(){
        return this.commandId;
    }


    public void info(){
        System.out.println(this.name);
        System.out.println(Arrays.toString(this.options));
        System.out.println(Arrays.toString(this.args));
        System.out.println();
    }
}
