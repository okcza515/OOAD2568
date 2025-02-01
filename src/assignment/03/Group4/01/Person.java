public class Person{
    public static void main(String[] args){
        IDevice LGSmartTelevision = new TV();
        IDevice PanasonicProjector = new Projector();
        IDevice AJHometheater = new SurroundSoundSystem();

        RemoteControl remote = new RemoteControl();

        remote.connectDevice(LGSmartTelevision);
        remote.turnOnDevice();
        remote.turnOffDevice();

        remote.connectDevice(PanasonicProjector);
        remote.turnOnDevice();
        remote.turnOffDevice();

        remote.connectDevice(AJHometheater);
        remote.turnOnDevice();
        remote.turnOffDevice();

    }
}