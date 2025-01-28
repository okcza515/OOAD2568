public class RemoteControl {
    IDevice device;


    
    public void turnoff() {
        device.turnOff();
    }

    public void turnon() {
        device.turnOn();
    }
}
