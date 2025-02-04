public class Person {

    public static void main(String[] args) {
        Projector projector = new Projector();
        TV tv = new TV();
        SurroundSoundSystem surroundSoundSystem = new SurroundSoundSystem();
        RemoteControl remote = new RemoteControl();

        remote.turnOn();
        remote.turnOff();


        remote = RemoteControl.connect(projector);
        remote.turnOn();
        remote.turnOff();

        remote = RemoteControl.connect(tv);
        remote.turnOn();
        remote.turnOff();

        remote = RemoteControl.connect(surroundSoundSystem);
        remote.turnOn();
        remote.turnOff();

    }
}
