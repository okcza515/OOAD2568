// Roodfan Maimahad 65070501049
public class RemoteControl {
    private static Device connectedDevice;
    
    // static
    public static void connectWithDevice(Device device) {
        connectedDevice = device;
    }

    public void turnOn() {
        connectedDevice.turnOn();
    }
    
    public void turnOff() {
        connectedDevice.turnOff();
    }
}
