public class TV implements IDevice {

    public void turnOn() {
        System.out.println("TV turn on...");
    }

    public void turnOff() {
        System.out.println("TV turn off...");
    }

    public String toString() {
        return "TV";
    }

}
