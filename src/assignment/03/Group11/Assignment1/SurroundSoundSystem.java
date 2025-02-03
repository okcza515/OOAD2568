// 65070503408 Jarukit Jintanasathirakul
class SurroundSoundSystem implements Device {

    @Override
    public void turnOn() {
        System.out.println("Sound System is now ON");
    }

    @Override
    public void turnOff() {
        System.out.println("Sound System is now OFF");
    }

    @Override
    public String getDeviceName() {
        return "Surround Sound System";
    }
}