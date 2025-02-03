//Pitchayuth Jampong 65070501079

public class Person {
    public static void main(String[] args) {
        RemoteControl remote = new RemoteControl();

        Device tv = new TV();
        Device projector = new Projector();
        Device surroundSoundSystem = new SurroundSoundSystem();

        RemoteControl.connectWithDevice(tv);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.connectWithDevice(projector);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.connectWithDevice(surroundSoundSystem);
        remote.turnOn();
        remote.turnOff();
    }
}