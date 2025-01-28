public class Projector implements IDevice {
    @Override
    public void turnOn() {
        System.out.println("Projector is now ON.");
    }

    @Override
    public void turnOff() {
        System.out.println("Projector is now OFF.");
    }
}