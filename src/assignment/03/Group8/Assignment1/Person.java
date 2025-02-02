//Thanoo Thanusuttiyaporn 65070503451
public class Person {
    
    public static void main(String[] args) {
        Device Projector = new Projector();
        Device TV = new TV();
        Device SurroundSoundSystem = new SurroundSoundSystem();

        RemoteControl remote = new RemoteControl();

        RemoteControl.pairingDevice(Projector);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.pairingDevice(TV);
        remote.turnOn();
        remote.turnOff();

        RemoteControl.pairingDevice(SurroundSoundSystem);
        remote.turnOn();
        remote.turnOff();
    }
}
