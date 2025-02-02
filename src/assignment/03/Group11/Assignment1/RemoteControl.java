public class RemoteControl {
    private Device currentDevice;

    public static void connectDevice(RemoteControl remoteControl, Device device) {
        remoteControl.currentDevice = device;
        System.out.println("Remote connected to: " + device.getDeviceName());
    }

    public void turnOn() {
        if (currentDevice != null) {
            currentDevice.turnOn();
        } else {
            System.out.println("No device connected");
        }
    }

    public void turnOff() {
        if (currentDevice != null) {
            currentDevice.turnOff();
        } else {
            System.out.println("No device connected");
        }
    }

}