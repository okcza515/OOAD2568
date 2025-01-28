public class Person {
    public static void main(String[] args) {
        IDevice projector = new Projector();
        IDevice tv = new TV();
        IDevice surroundSoundSystem = new SurroundSoundSystem();

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
