public class RemoteControl {

    private static Device pairedDevice;

    public static void pairingDevice(Device device) {
        pairedDevice = device;
    }

    public void turnOn() {
       if (pairedDevice != null) {
            pairedDevice.turnOn();
            System.out.println("Remote connected to: " + pairedDevice.getClass().getSimpleName());
        } else {
            System.out.println("No device paired.");
        }
    }

    public void turnOff() {
        if (pairedDevice != null) {
            pairedDevice.turnOff();
            System.out.println("Remote disconnected from: " + pairedDevice.getClass().getSimpleName());
        } else {
            System.out.println("No device paired.");
        }
    }
}

//Supanut Wongtanom 65070503437