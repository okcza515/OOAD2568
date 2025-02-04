public class RemoteControl {
    private Device connectedDevice;

    public void pairedDevice(Device device) {
        this.connectedDevice = device;
        System.out.println("Connected");
    }

    public void turnDeviceOn() {
       this.connectedDevice.turnDeviceOn();
    }

    public void turnDeviceOff() {
        this.connectedDevice.turnDeviceOff();
    }
}