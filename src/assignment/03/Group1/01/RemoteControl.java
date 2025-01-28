public class RemoteControl {
    IDevice device;

    static public RemoteControl connectDevice(IDevice connectedDevice) {
        RemoteControl remote = new RemoteControl();
        remote.device = connectedDevice;
        System.out.println("Connect to device " + remote.device);
        return remote;
    };
    
    public void turnOff() {
        device.turnOff();
    }

    public void turnOn() {
        device.turnOn();
    }
}
