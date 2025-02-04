class SurroundSoundSystem implements Device {

    @Override
    public void turnDeviceOn() {
        System.out.println("Surround Sound System is turn ON");
    }

    @Override
    public void turnDeviceOff() {
        System.out.println("Surround Sound System is turn OFF");
    }

    public String getDeviceName() {
        return "Surround Sound System";
    }
}