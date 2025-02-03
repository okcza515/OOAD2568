// 65070503408 Jarukit Jintanasathirakul
class TV implements Device {

    @Override
    public void turnOn() {
        System.out.println("TV is now ON");
    }

    @Override
    public void turnOff() {
        System.out.println("TV is now OFF");
    }

    @Override
    public String getDeviceName() {
        return "TV";
    }
}