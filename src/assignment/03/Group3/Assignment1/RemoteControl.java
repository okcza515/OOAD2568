
// 65070501001 Kantapong Vongpanich
public class RemoteControl {

    private static Device selectedDevice;

    public static void chooseDevice(Device device) {
        selectedDevice = device;
    }

    public void turnOn() {
        selectedDevice.turnOn();
    }

    public void turnOff() {
        selectedDevice.turnOff();
    }
}
