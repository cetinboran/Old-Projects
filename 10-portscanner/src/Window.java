import java.util.Objects;
import java.util.Scanner;

public class Window {
    private String ip = " ", port = " ";
    int[] ports;
    public void loop(){
        boolean over = false;
        String command;

        System.out.println("-h / --help for Help");
        while(!over){
            Scanner input = new Scanner(System.in);
            System.out.print(">: ");

            command = input.nextLine().toLowerCase();

            if(Objects.equals(command, "-q") || Objects.equals(command, "--quit")) over = true;
            else if(Objects.equals(command, "-h") || Objects.equals(command, "--help")) help();
            else if(Objects.equals(command, "--show")) show();
            else if(Objects.equals(command, "run")){
                if(!Objects.equals(this.ip, " ") && this.ports != null){
                    break;
                }
                System.out.println("You have to set the host and ports");
            }
            else{
                handleArgs(command);
            }
        }

        if(over) System.exit(0);
    }

    private void help(){
        System.out.println("You can set variables by using 'SET' keyword");
        System.out.println("Ex: set ip/port 8.8.8.8");
        System.out.println("You can use '--SHOW' to see all your variables");
        System.out.println("You can use '-' for range of ports or just one or two with ','");
        System.out.println("Ex: set port 1-23. Ex: set port 21,22,80,443,53");
        System.out.println("-q / --quit for quiting abv.");
    }

    private void show(){
        System.out.printf("HOST: %s\nPorts: %s\n", this.ip, this.port);
    }

    private void handleArgs(String command){
        String[] args = command.split(" ");

        for(int i = 0; i < args.length; i++){
            args[i] = args[i].trim();
        }

        if(args.length < 3){
            System.out.println("Invalid Argument!");
            return;
        }

        if(!Objects.equals(args[0], "set") || (!Objects.equals(args[1], "ip") && !Objects.equals(args[1], "port"))){
            System.out.println("Invalid Argument!");
            return;
        }

        if(Objects.equals(args[1], "ip")){
            this.ip = args[2];
        }
        else{
            this.port = args[2];
            this.ports = handlePorts(args[2]);
        }

    }

    private int[] handlePorts(String ports){
        int[] intPorts = null;
        if(ports.contains("-")){
            String[] portsStr = ports.split("-");
            int start = Integer.parseInt(portsStr[0]);
            int end = Integer.parseInt(portsStr[1]);

            if(start > end){
                System.out.println("Invalid Port");
                return new int[] {-1};
            }

            intPorts = new int[end - start];

            int i = 0;
            while(start < end){
                intPorts[i++] = start++;
            }
        } else if (ports.contains(",")) {
            String[] portsStr = ports.split(",");
            intPorts = new int[portsStr.length];

            int k = 0;
            for(String x: portsStr){
                intPorts[k++] = Integer.parseInt(x);
            }
        }


        return intPorts;
    }

    public int[] getPorts(){
        return this.ports;
    }

    public String getHost(){
        return this.ip;
    }
}
