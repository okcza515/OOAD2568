public class Person {
    public static void main(String[] args) {
        Device projector = new Projector();
        Device tv = new TV();
        Device surround_sound_system = new SurroundSoundSystem();

        RemoteControl remote = new RemoteControl();

        RemoteControl.chooseDevice(projector);
        remote.turnOn();
        remote.turnOff();
    
        RemoteControl.chooseDevice(tv);
        remote.turnOn();
        remote.turnOff();
        
        RemoteControl.chooseDevice(surround_sound_system);
        remote.turnOn();
        remote.turnOff();
    }
}

// 65070501039 Pongpon Butseemart