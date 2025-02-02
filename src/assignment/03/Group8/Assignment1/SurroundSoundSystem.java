public class SurroundSoundSystem implements Control{
    @Override
    public void turnOn(){
        System.out.println("Surround Sound System is turned on");
    }
    @Override
    public void turnOff(){
        System.out.println("Surround Sound System is turned off");
    }
    public void adjustVolume(){
        System.out.println("Adjusting volume");
    }
    @Override
    public void special(){
        adjustVolume();
    }
}
//Sinsorn Chaithavornkit 65070503436