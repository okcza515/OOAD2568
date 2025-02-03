public class RemoteControl {
    private Device currentDevice;

    public void pairDevice(Device device) {
        this.currentDevice = device;
        System.out.println("RemoteControl paired with " + device.getClass().getSimpleName() + ".");
    }

    public void turnOn() {
        if (currentDevice != null) {
            currentDevice.turnOn();
        } else {
            System.out.println("No device paired.");
        }
    }

    public void turnOff() {
        if (currentDevice != null) {
            currentDevice.turnOff();
        } else {
            System.out.println("No device paired.");
        }
    }
}

//65070501042 Pakaporn Kanteng