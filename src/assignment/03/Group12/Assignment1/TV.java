public class TV implements Device {
    @Override
    public void turnDeviceOn() {
        System.out.println("TV device is turn on");
    }

    @Override
    public void turnDeviceOff() {
        System.out.println("TV device is turn off");
    }
}