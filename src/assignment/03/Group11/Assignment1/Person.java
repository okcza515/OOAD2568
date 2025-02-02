public class Person {
    public static void main(String[] args) {
        RemoteControl remote = new RemoteControl();

        Device projector = new Projector();
        RemoteControl.connectDevice(remote, projector);
        remote.turnOn();
        remote.turnOff();
    }
}