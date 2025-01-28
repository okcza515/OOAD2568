public class Person {
    public static void main(String[] args) {
        IDevice projector = new Projector();
        IDevice tv = new TV();
        IDevice surroundSoundSystem = new SurroundSoundSystem();

        RemoteControl.connectDevice(projector);
        RemoteControl remote = new RemoteControl();
        remote.turnOn();
        remote.turnOff();

        RemoteControl.connectDevice(tv);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.connectDevice(surroundSoundSystem);
        remote.turnOn();
        remote.turnOff();
    }
}
