//Natcha Trairattanasak 65070503415
public class Projector implements Control {
    @Override
    public void turnOn(){
        System.out.println("Projector is turned on");
    }
    @Override
    public void turnOff(){
        System.out.println("Projector is turned off");
    }
    public void adjustScreen(){
        System.out.println("Performing auto screen adjustment");
    }
    @Override
    public void special(){
        adjustScreen();
    }
}
