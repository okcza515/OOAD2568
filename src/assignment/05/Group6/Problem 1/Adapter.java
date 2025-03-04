//65070501008
public class Adapter extends RoundPeg{
    private SquarePeg peg;

    public Adapter(SquarePeg peg){
        this.peg = peg;
    }

    @Override
    public double getRadius(){
        return (Math.sqrt(Math.pow((peg.getWidth() / 2), 2) * 2));
    }
}
