public class Person {
    public static void main(String[] args) {

        TV tv = new TV();
        Projector projector = new Projector();
        SurroundSoundSystem soundSystem = new SurroundSoundSystem();

        RemoteControl myRemote = new RemoteControl();

        myRemote.pairDevice(tv);
        myRemote.turnOn();
        myRemote.turnOff();

        myRemote.pairDevice(projector);
        myRemote.turnOn();
        myRemote.turnOff();

        myRemote.pairDevice(soundSystem);
        myRemote.turnOn();
        myRemote.turnOff();
    }
}

//65070501042 Pakaporn Kanteng