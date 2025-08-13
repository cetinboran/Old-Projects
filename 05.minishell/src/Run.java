import Commands.Echo;
import Commands.Wc;
import Tokens.Pipe;

import java.util.Arrays;

public class Run {
    // Burada echo wc gibi commandleri tutup çalıştırma işlemleri yapacağız.
    private Object[] commands;

    public Run(String text){
        new Classifier(text).Classify();
        commands = Classifier.getObject();
    }

    public void run(){
        //System.out.println(Arrays.toString(commands));


        for(Object x: commands){
            if(x == null)
                continue;

            if(x.getClass() == Echo.class){
                ((Echo) x).run();
            }
            else if(x.getClass() == Pipe.class){
                ((Pipe) x).transferOutput();
                // Burada transfer işlemini yap çıktılar ile
            }
            else if(x.getClass() == Wc.class){
                ((Wc) x).run();
            }
            /*
            else if(x.getClass() == Pipe.class){
                System.out.println(((Pipe) x).getLeft());
                System.out.println(((Pipe) x).getRight());
            }
             */
        }
    }


    public Object[] getCommands() {
        return commands;
    }

    public void setCommands(Object[] commands){
        this.commands = commands;
    }

}
