import java.io.IOException;
import java.net.Socket;

public class Main extends Thread {
    private String ip;
    private int port;
    public Main(String ip, int port){
        this.ip = ip;
        this.port = port;
    }
    public static void main(String[] args) {
        Window window = new Window();
        window.loop();

        String host = window.getHost();
        int[] ports = window.getPorts();

        for(int port: ports){
            Main scan = new Main(host, port);
            scan.start();
        }
    }

    public void run(){
        connection(this.ip, this.port);
    }
    public void connection(String ip, int port){

        Socket socket = null;
        try{
            socket = new Socket(ip, port);

            System.out.printf("%d: Open\n", port);

            socket.shutdownInput();
            socket.shutdownOutput();
            socket.close();
        }
        catch (IOException e){
            System.out.printf("%d: Close|Filtered\n", port);
        }

    }
}
