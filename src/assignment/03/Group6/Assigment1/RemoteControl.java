public class RemoteControl {

    private IControllDevice device;

    public static RemoteControl connect(IControllDevice device) {
        RemoteControl remote = new RemoteControl();
        remote.device = device;
        System.out.println("Remote control has been connected to: " + device.getClass().getSimpleName() + ".");
        return remote;
    }

    public void turnOn(){
        if(device != null){
            device.turnOn();
        }else{
            System.out.println("Pleas connect to a device first.");
        }
    }

    public void turnOff(){
        if(device != null){
            device.turnOff();
        }else{
            System.out.println("Pleas connect to a device first.");
        }
    }
}
