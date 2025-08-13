package Tokens;

import Commands.Command;
import Commands.Echo;
import Commands.Wc;

public class Pipe {
    private final int pipeId;

    // Buradaki right ve left bizim | in sağ ve solundaki komutları tutuyor. Buradan input transfer at.
    Object right, left;

    public Pipe(Object left, Object right){
        this.pipeId = 0;

        this.right = right;
        this.left = left;

        setStdO();
        setStdI();
    }

    public void transferOutput(){
        // Command classının özellikleri olduğu için böyle çağırabildim.
        if(left.getClass() != null && right.getClass() != null){
            String output = ((Command) left).getOutput();
            ((Command) right).setInput(output);
        }
    }

    private void setStdO(){
        // pipe'ın solunda stdO'ları 1 yapıyorum ki çıktıları pipe'a yollasın.
        if(left.getClass() != null && left.getClass() != Pipe.class){
            ((Command) left).setStdO(true);
        }
    }

    private void setStdI(){
        if(right.getClass() != null && right.getClass() != Pipe.class){
            ((Command) right).setStdI(true);
        }
    }

    public Object getRight() {
        return right;
    }
    public void setRight(Object right) {
        this.right = right;
    }

    public Object getLeft() {
        return left;
    }

    public void setLeft(Object left) {
        this.left = left;
    }

    public int getId(){
        return this.pipeId;
    }
}
