class SurroundSoundSystem implements Device {
    @Override
    public void turnOn() {
        System.out.println("Surround Sound System is turned on.");
    }

    @Override
    public void turnOff() {
        System.out.println("Surround Sound System is turned off.");
    }
}