public class Adapter extends RoundPeg{
    private SquarePeg squarepeg;

    public Adapter(SquarePeg squarepeg){
        this.squarepeg = squarepeg;
    }
    
    @Override
    public double getRadius(){
        return (Math.sqrt(Math.pow((squarepeg.getWidth() / 2), 2) * 2));
    }
}
//Korawit Sritotum 65070503402