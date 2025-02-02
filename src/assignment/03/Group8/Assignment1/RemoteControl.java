//Pontakorn Wichaporn 65070503427
public class RemoteControl {
    private static Control choosedDevice;

    public static void pairingDevice(Control device){
        choosedDevice = device;
    }

    public void turnOn(){
        choosedDevice.turnOn();
    }

    public void turnOff(){
        choosedDevice.turnOff();
    }

}
