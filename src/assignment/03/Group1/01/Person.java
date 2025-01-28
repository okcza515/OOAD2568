public class Person {
    public static void main(String[] args) {
        Device projector = new Projector();
        Device tv = new TV();
        Device surroundSoundSystem = new SurroundSoundSystem();

        RemoteControl.chooseDevice(projector);
        RemoteControl remote = new RemoteControl();
        remote.turnOn();
        remote.turnOff();

        RemoteControl.chooseDevice(tv);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.chooseDevice(surroundSoundSystem);
        remote.turnOn();
        remote.turnOff();
    }
}
