public class Person {
    public static void main(String[] args) {
        RemoteControl remote = new RemoteControl();

        TV tv = new TV();
        remote.pairedDevice(tv);
        remote.turnDeviceOn();
        remote.turnDeviceOff();

        Projector projector = new Projector();
        remote.pairedDevice(projector);
        remote.turnDeviceOn();
        remote.turnDeviceOff();
    }
}