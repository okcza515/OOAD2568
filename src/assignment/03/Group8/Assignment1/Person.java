//Thanoo Thanusuttiyaporn 65070503451
public class Person {
    
    public static void main(String[] args) {
        Projector Projector = new Projector();
        Control TV = new TV();
        Control SurroundSoundSystem = new SurroundSoundSystem();

        RemoteControl remote = new RemoteControl();

        RemoteControl.pairingDevice(Projector);
        remote.turnOn();
        Projector.adjustScreen();
        remote.turnOff();
        // remote.special();

        RemoteControl.pairingDevice(TV);
        remote.turnOn();
        remote.turnOff();
        remote.special();

        RemoteControl.pairingDevice(SurroundSoundSystem);
        remote.turnOn();
        remote.turnOff();
        remote.special();
    }
}
