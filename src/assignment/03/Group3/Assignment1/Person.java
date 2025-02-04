
public class Person {

    public static void main(String[] args) {
        Device projector = new Projector();
        Device tv = new TV();
        Device surround_sound_system = new SurroundSoundSystem();

        RemoteControl.getInstance().chooseDevice(projector);
        RemoteControl.getInstance().turnOn();
        RemoteControl.getInstance().turnOff();

        RemoteControl.getInstance().chooseDevice(tv);
        RemoteControl.getInstance().turnOn();
        RemoteControl.getInstance().turnOff();

        RemoteControl.getInstance().chooseDevice(surround_sound_system);
        RemoteControl.getInstance().turnOn();
        RemoteControl.getInstance().turnOff();
    }
}

// 65070501039 Pongpon Butseemart
