import Commands.Echo;
import Commands.Wc;
import Tokens.Pipe;

import java.lang.reflect.Array;
import java.util.Arrays;
import java.util.Objects;

public class Classifier {
    private final String text;
    private static Object[] commands;
    public Classifier(String text){
        this.text = this.simplerFormat(text);
    }

    public void Classify(){
        // Burada | a bölerek komut sayısını öğreniyoruz
        String[] byPipe = this.text.split("\\|");

        // komut sayısı kadar obje arrayi oluşturduk
        commands = new Object[2*byPipe.length - 1];
        Object[] justCommands = new Object[byPipe.length];

        int i = 0;
        while(i < byPipe.length){
            String[] simple = byPipe[i].trim().split(" ");

            String command = simple[0].trim();
            String options = "";
            String arguments = "";

            int j = 1;
            while(j < simple.length){


                if(simple[j].trim().charAt(0) == '-'){
                    options += simple[j].trim() + " ";
                }
                else{
                    arguments += simple[j] + " ";
                }
                j++;
            }

            // burada komutun türünü belirleyip onu static arrayimize atyıyoruz.
            checkCommands(command, options.trim(), arguments, i, justCommands);
            i++;
        }

        handlePipes(justCommands, commands);
    }

    private void checkCommands(String command, String options, String args, int index, Object[] commands){
        // Komutları sınıflandırıyoruz ve değişkene yüklüyoruz.

        if(Objects.equals(command, "echo")){
            Echo x = new Echo(command, args, options);
            commands[index] = x;
        }
        else if(Objects.equals(command, "wc")){
            Wc x = new Wc(command, args, options);
            commands[index] = x;
        }
    }

    private void handlePipes(Object[] justCommands, Object[] commands){
        for(int i = 0; i < justCommands.length; i++){
            commands[i*2] = justCommands[i];
        }

        for(int i = 0; i < commands.length; i++){
            if(commands[i] == null){
                Pipe x = new Pipe(commands[i - 1], commands[i + 1]);
                commands[i] = x;
            }
        }
    }

    private int passPr(String text, int i){
        // Gelen text'te her kelime arasını 1 boşluk yapıyor.
        for(int j = i; j < text.length(); j++){
            char c = text.charAt(i);
            if(c == ' '){
                i++;
            }
        }

        return --i;
    }

    private String simplerFormat(String text){
        // Daha rahat uğraşılabilir formata getirdik text'i
        int i = 0;
        StringBuilder newText = new StringBuilder();

        while(i < text.length()){
            char c = text.charAt(i);
            if(c == ' '){
                i = this.passPr(text, i);
            }
            newText.append(c);
            i++;
        }

        return newText.toString();
    }



    public void listObject(){
        System.out.println(Arrays.toString(commands));
    }

    public static Object[] getObject(){
        return commands;
    }
}
