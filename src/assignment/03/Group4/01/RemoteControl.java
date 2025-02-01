public class RemoteControl {
    private static IDevice _connectedDevice;

    public static void connectDevice(IDevice device) {
        _connectedDevice = device;
        System.out.println("Device is connected successfully");
    }

    public void turnOnDevice() {
        if (_connectedDevice != null) {
            _connectedDevice.turnOn();
        } else {
            System.out.println("Device is not connected");
        }
    }

    public void turnOffDevice() {
        if (_connectedDevice != null) {
            _connectedDevice.turnOff();
        } else {
            System.out.println("Device is not connected");
        }
    }
}