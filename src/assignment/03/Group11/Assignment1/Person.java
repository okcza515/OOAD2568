public class Person {
    public static void main(String[] args) {
        RemoteControl remote = new RemoteControl();

        Device projector = new Projector();
        Device surroundSoundSystem = new SurroundSoundSystem();
        Device TV = new TV();

        RemoteControl.connectDevice(remote, projector);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.connectDevice(remote, surroundSoundSystem);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.connectDevice(remote, TV);
        remote.turnOn();
        remote.turnOff();


    }
}