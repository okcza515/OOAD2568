public class Person {
    public static void main(String[] args) {
        Projector projector = new Projector();
        TV tv = new TV();
        SurroundSoundSystem surroundSoundSystem = new SurroundSoundSystem();

        RemoteControl remote = new RemoteControl();

        RemoteControl.connectToDevice(projector);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.connectToDevice(tv);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.connectToDevice(surroundSoundSystem);
        remote.turnOn();
        remote.turnOff();
    }
}