class RemoteControl {
    private static Device connectedDevice;

    public static void connectToDevice(Device device) {
        connectedDevice = device;
        System.out.println("Remote connected to:" + device.getClass().getSimpleName() );
    }

    public void turnOn() {
        if (connectedDevice != null) {
            connectedDevice.turnOn();
        } else {
            System.out.println("No device connected.");
        }
    }

    public void turnOff() {
        if (connectedDevice != null) {
            connectedDevice.turnOff();
        } else {
            System.out.println("No device connected.");
        }
    }
}