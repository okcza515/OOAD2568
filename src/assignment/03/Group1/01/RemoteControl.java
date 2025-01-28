public class RemoteControl {
    IDevice device;

    static public RemoteControl connectDevice(IDevice connected_device) {
        RemoteControl remote = new RemoteControl();
        remote.device = connected_device;
        System.out.println("Connect to device " + remote.device);
        return remote;
    };
    
    public void turnoff() {
        device.turnOff();
    }

    public void turnon() {
        device.turnOn();
    }
}
