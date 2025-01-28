public class SurroundSoundSystem implements IDevice{
    @Override
    public void turnOn() {
        System.out.println("Surround sound system is now On.");
    }
    @Override
    public void turnOff() {
        System.out.println("Surround sound system is now Off.");
    }
}