
// 65070501001 Kantapong Vongpanich
public class RemoteControl {

    private static Device selectedDevice;
    private static final RemoteControl instance = new RemoteControl();

    public static RemoteControl getInstance() {
        return instance;
    }

    public void chooseDevice(Device device) {
        selectedDevice = device;
    }

    public void turnOn() {
        selectedDevice.turnOn();
    }

    public void turnOff() {
        selectedDevice.turnOff();
    }
}
