public class Projector implements Device {
    @Override
    public void turnDeviceOn() {
        System.out.println("Projector lights on");
    }

    @Override
    public void turnDeviceOff() {
        System.out.println("Projector lights off");
    }
}