public class Person{
    public static void main(String[] args) {
        Device tv = new TV();
        Device projector = new Projectors();
        Device soundsystem = new SurroundSoundSystems();

        RemoteControls remote = new RemoteControls();

        RemoteControls.pairingDevice(tv);
        remote.turnOn();
        remote.turnOff();
        System.out.println("Connected to TV");

        RemoteControls.pairingDevice(soundsystem);
        remote.turnOn();
        remote.turnOff();
        System.out.println("Connected to Soundsystem");

        RemoteControls.pairingDevice(projector);
        remote.turnOn();
        remote.turnOff();
        System.out.println("Connected to Projector");

    }
}
//Korawit Sritotum 65070503402