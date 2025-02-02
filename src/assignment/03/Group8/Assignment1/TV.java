public class TV implements Control {
    @Override
    public void turnOn(){
        System.out.println("Turn on the TV");
    }
    @Override
    public void turnOff(){
        System.out.println("Turn off the TV");
    }
    public void playMovie(){
        System.out.println("A movie is played on the TV");
    }
    @Override
    public void special(){
        playMovie();
    }
}

//Kittipob Borisut 65070503407