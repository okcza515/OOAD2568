public class Person{
    public static void main(String[] args) {
        Device tv = new TV();
        Device projector = new Projector();
        Device soundsystem = new SurroundSoundSystem();

        RemoteControl remote = new RemoteControl();

        RemoteControl.pairingDevice(tv);
        remote.turnOn();
        remote.turnOff();
        System.out.println("Connected to TV");

        RemoteControl.pairingDevice(soundsystem);
        remote.turnOn();
        remote.turnOff();
        System.out.println("Connected to Soundsystem");

        RemoteControl.pairingDevice(projector);
        remote.turnOn();
        remote.turnOff();
        System.out.println("Connected to Projector");

    }
}
//Korawit Sritotum 65070503402