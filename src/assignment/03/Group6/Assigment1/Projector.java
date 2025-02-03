public class Projector implements IControllDevice {

    @Override
    public void turnOn(){
        System.out.println("Turn on projector.");
    };

    @Override
    public void turnOff(){
        System.out.println("Turn off projector.");
    }

}
